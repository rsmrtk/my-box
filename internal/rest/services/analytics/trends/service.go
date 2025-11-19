package trends

import (
	"context"
	"fmt"
	"time"

	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx context.Context
	req *da.TrendsRequest
	f   *Facade

	trends []da.MonthlyTrend
}

func (s *service) analyze() error {
	// Fetch trend data
	if err := s.fetchTrendData(); err != nil {
		return err
	}

	// Calculate trend changes after fetching data
	s.calculateTrendChanges()

	return nil
}

func (s *service) fetchTrendData() error {
	months := s.req.Months
	if months <= 0 {
		months = 6
	}

	query := fmt.Sprintf(`
		SELECT
			DATE_TRUNC('month', expense_date)::DATE as month,
			SUM(expense_amount) as total,
			COUNT(*) as transactions
		FROM expenses
		WHERE expense_date >= CURRENT_DATE - INTERVAL '%d months'
		GROUP BY month
		ORDER BY month ASC
	`, months)

	rows, err := s.f.pkg.M.FinDash.DB.Query(s.ctx, query)
	if err != nil {
		return errs.FailedToGetTrends
	}
	defer rows.Close()

	var trends []da.MonthlyTrend

	for rows.Next() {
		var trend da.MonthlyTrend
		var total float64
		var monthDate time.Time

		err := rows.Scan(&monthDate, &total, &trend.Count)
		if err != nil {
			return errs.FailedToScanData
		}

		// Get currency from request or use default
		currency := s.req.Currency
		if currency == "" {
			currency = "USD"
		}
		currencySymbol := da.GetCurrencySymbol(currency)

		trend.Month = models.NewDate(monthDate)
		trend.Total = []*models.Amount{{
			Amount:         total,
			CurrencyCode:   currency,
			CurrencySymbol: currencySymbol,
		}}

		trends = append(trends, trend)
	}

	s.trends = trends
	return nil
}

func (s *service) calculateTrendChanges() {
	// Get currency from request or use default
	currency := s.req.Currency
	if currency == "" {
		currency = "USD"
	}
	currencySymbol := da.GetCurrencySymbol(currency)

	// Calculate month-over-month changes
	for i := 1; i < len(s.trends); i++ {
		currentTotal := s.trends[i].Total[0].Amount
		previousTotal := s.trends[i-1].Total[0].Amount

		// Calculate absolute change
		change := currentTotal - previousTotal
		s.trends[i].Change = []*models.Amount{{
			Amount:         change,
			CurrencyCode:   currency,
			CurrencySymbol: currencySymbol,
		}}

		// Calculate percentage change
		if previousTotal > 0 {
			s.trends[i].ChangePercentage = (change / previousTotal) * 100
		} else if currentTotal > 0 {
			s.trends[i].ChangePercentage = 100
		}
	}
}

func (s *service) reply() *da.TrendsResponse {
	return &da.TrendsResponse{
		Trends: s.trends,
	}
}
