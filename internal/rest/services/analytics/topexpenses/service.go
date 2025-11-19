package topexpenses

import (
	"context"

	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx context.Context
	req *da.TopExpensesRequest
	f   *Facade

	categories []da.CategorySummary
}

func (s *service) find() error {
	// Fetch top expenses data
	if err := s.fetchTopExpenses(); err != nil {
		return err
	}

	// Calculate percentages after fetching data
	s.calculatePercentages()

	return nil
}

func (s *service) fetchTopExpenses() error {
	limit := s.req.Limit
	if limit <= 0 {
		limit = 3
	}

	query := `
		SELECT
			expense_type,
			SUM(expense_amount) as total,
			COUNT(*) as count
		FROM expenses
		WHERE expense_type IS NOT NULL
		GROUP BY expense_type
		ORDER BY total DESC
		LIMIT $1
	`

	rows, err := s.f.pkg.M.FinDash.DB.Query(s.ctx, query, limit)
	if err != nil {
		return errs.FailedToGetTopExpenses
	}
	defer rows.Close()

	var categories []da.CategorySummary

	// Fetch data from database
	for rows.Next() {
		var cat da.CategorySummary
		var total float64

		err := rows.Scan(&cat.Category, &total, &cat.Count)
		if err != nil {
			return errs.FailedToScanData
		}

		// Get currency from request or use default
		currency := s.req.Currency
		if currency == "" {
			currency = "USD"
		}
		currencySymbol := da.GetCurrencySymbol(currency)

		// Create Amount object
		cat.Total = []*models.Amount{{
			Amount:         total,
			CurrencyCode:   currency,
			CurrencySymbol: currencySymbol,
		}}

		categories = append(categories, cat)
	}

	s.categories = categories
	return nil
}

func (s *service) calculatePercentages() {
	// Calculate total across all categories
	var totalAllCategories float64
	for _, cat := range s.categories {
		if len(cat.Total) > 0 {
			totalAllCategories += cat.Total[0].Amount
		}
	}

	// Calculate percentage for each category
	for i := range s.categories {
		if totalAllCategories > 0 && len(s.categories[i].Total) > 0 {
			s.categories[i].Percentage = (s.categories[i].Total[0].Amount / totalAllCategories) * 100
		}
	}
}

func (s *service) reply() *da.TopExpensesResponse {
	return &da.TopExpensesResponse{
		Categories: s.categories,
	}
}
