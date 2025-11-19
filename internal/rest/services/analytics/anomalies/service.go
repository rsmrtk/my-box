package anomalies

import (
	"context"
	"time"

	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx context.Context
	req *da.AnomaliesRequest
	f   *Facade

	categoryAverages map[string]float64
	expenses         []expenseData
	anomalies        []da.AnomalyItem
}

type expenseData struct {
	ID     string
	Name   string
	Amount float64
	Type   string
	Date   time.Time
}

func (s *service) detect() error {
	// Fetch category averages
	if err := s.fetchCategoryAverages(); err != nil {
		return err
	}

	// Fetch recent expenses
	if err := s.fetchRecentExpenses(); err != nil {
		return err
	}

	// Detect anomalies based on fetched data
	s.detectAnomalies()

	return nil
}

func (s *service) fetchCategoryAverages() error {
	query := `
		SELECT expense_type, AVG(expense_amount)
		FROM expense
		WHERE expense_date >= CURRENT_DATE - INTERVAL '3 months'
		  AND expense_type IS NOT NULL
		GROUP BY expense_type
	`

	rows, err := s.f.pkg.M.FinDash.DB.Query(s.ctx, query)
	if err != nil {
		return errs.FailedToCalculateAverages
	}
	defer rows.Close()

	avgMap := make(map[string]float64)
	for rows.Next() {
		var category string
		var avg float64
		err := rows.Scan(&category, &avg)
		if err != nil {
			return errs.FailedToScanData
		}
		avgMap[category] = avg
	}

	s.categoryAverages = avgMap
	return nil
}

func (s *service) fetchRecentExpenses() error {
	query := `
		SELECT
			expense_id,
			expense_name,
			expense_amount,
			expense_type,
			expense_date
		FROM expense
		WHERE expense_date >= CURRENT_DATE - INTERVAL '1 month'
		  AND expense_type IS NOT NULL
		ORDER BY expense_amount DESC
	`

	rows, err := s.f.pkg.M.FinDash.DB.Query(s.ctx, query)
	if err != nil {
		return errs.FailedToFindAnomalies
	}
	defer rows.Close()

	var expenses []expenseData
	for rows.Next() {
		var exp expenseData
		err := rows.Scan(&exp.ID, &exp.Name, &exp.Amount, &exp.Type, &exp.Date)
		if err != nil {
			return errs.FailedToScanData
		}
		expenses = append(expenses, exp)
	}

	s.expenses = expenses
	return nil
}

func (s *service) detectAnomalies() {
	threshold := s.req.Threshold
	if threshold <= 0 {
		threshold = 1.5
	}

	var anomalies []da.AnomalyItem

	for _, exp := range s.expenses {
		// Check if category average exists
		if avg, ok := s.categoryAverages[exp.Type]; ok {
			deviationFactor := exp.Amount / avg

			// Check if it's an anomaly based on threshold
			if deviationFactor > threshold {
				item := da.AnomalyItem{
					ID:              exp.ID,
					Name:            exp.Name,
					Type:            exp.Type,
					Date:            models.NewDate(exp.Date),
					DeviationFactor: deviationFactor,
					Amount: []*models.Amount{{
						Amount:         exp.Amount,
						CurrencyCode:   "USD",
						CurrencySymbol: "$",
					}},
					CategoryAverage: []*models.Amount{{
						Amount:         avg,
						CurrencyCode:   "USD",
						CurrencySymbol: "$",
					}},
				}

				// Determine anomaly severity level
				switch {
				case deviationFactor > 3:
					item.Status = "Critical"
				case deviationFactor > 2:
					item.Status = "High"
				default:
					item.Status = "Medium"
				}

				anomalies = append(anomalies, item)
			}
		}
	}

	s.anomalies = anomalies
}

func (s *service) reply() *da.AnomaliesResponse {
	return &da.AnomaliesResponse{
		Anomalies: s.anomalies,
	}
}
