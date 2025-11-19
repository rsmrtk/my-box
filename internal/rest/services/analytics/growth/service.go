package growth

import (
	"context"

	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx context.Context
	req *da.GrowthRequest
	f   *Facade

	currentMonth     float64
	previousMonth    float64
	growthAmount     float64
	growthPercentage float64
}

func (s *service) calculate() error {
	// Fetch income data
	if err := s.fetchIncomeData(); err != nil {
		return err
	}

	// Calculate growth metrics after fetching data
	s.calculateGrowthMetrics()

	return nil
}

func (s *service) fetchIncomeData() error {
	// Current month income
	err := s.f.pkg.M.FinDash.DB.QueryRow(s.ctx, `
		SELECT COALESCE(SUM(income_amount), 0)
		FROM income
		WHERE DATE_TRUNC('month', income_date) = DATE_TRUNC('month', CURRENT_DATE)
	`).Scan(&s.currentMonth)
	if err != nil {
		return errs.FailedToCalculateGrowth
	}

	// Previous month income
	err = s.f.pkg.M.FinDash.DB.QueryRow(s.ctx, `
		SELECT COALESCE(SUM(income_amount), 0)
		FROM income
		WHERE DATE_TRUNC('month', income_date) = DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
	`).Scan(&s.previousMonth)
	if err != nil {
		return errs.FailedToCalculateGrowth
	}

	return nil
}

func (s *service) calculateGrowthMetrics() {
	// Calculate growth amount
	s.growthAmount = s.currentMonth - s.previousMonth

	// Calculate growth percentage
	if s.previousMonth > 0 {
		s.growthPercentage = ((s.currentMonth - s.previousMonth) / s.previousMonth) * 100
	} else if s.currentMonth > 0 {
		// If previous month is 0 but current month has income, show 100% growth
		s.growthPercentage = 100
	} else {
		// Both are 0, no growth
		s.growthPercentage = 0
	}
}

func (s *service) reply() *da.GrowthResponse {
	return &da.GrowthResponse{
		CurrentMonth: []*models.Amount{{
			Amount:         s.currentMonth,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		PreviousMonth: []*models.Amount{{
			Amount:         s.previousMonth,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		GrowthAmount: []*models.Amount{{
			Amount:         s.growthAmount,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		GrowthPercentage: s.growthPercentage,
	}
}
