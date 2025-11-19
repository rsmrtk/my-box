package dashboard

import (
	"context"

	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx context.Context
	req *da.DashboardRequest
	f   *Facade

	// Income data
	totalIncome    float64
	monthlyIncome  float64
	dailyAvgIncome float64

	// Expense data
	totalExpense    float64
	monthlyExpense  float64
	dailyAvgExpense float64

	// Calculated metrics
	netCashFlow     float64
	savingsRate     float64
	stabilityRatio  float64
	stabilityStatus string

	// Categories
	topCategories []da.CategorySummary
}

func (s *service) calculate() error {
	// Calculate income metrics
	if err := s.calculateIncomeMetrics(); err != nil {
		return err
	}

	// Calculate expense metrics
	if err := s.calculateExpenseMetrics(); err != nil {
		return err
	}

	// Get top expense categories
	if err := s.getTopCategories(); err != nil {
		return err
	}

	// Calculate cash flow metrics after fetching all data
	s.calculateCashFlowMetrics()

	return nil
}

func (s *service) calculateIncomeMetrics() error {
	// Total income
	err := s.f.pkg.M.FinDash.DB.QueryRow(s.ctx, `
		SELECT COALESCE(SUM(income_amount), 0)
		FROM income
	`).Scan(&s.totalIncome)
	if err != nil {
		return errs.FailedToCalculateIncome
	}

	// Monthly income (current month)
	err = s.f.pkg.M.FinDash.DB.QueryRow(s.ctx, `
		SELECT COALESCE(SUM(income_amount), 0)
		FROM income
		WHERE DATE_TRUNC('month', income_date) = DATE_TRUNC('month', CURRENT_DATE)
	`).Scan(&s.monthlyIncome)
	if err != nil {
		return errs.FailedToCalculateIncome
	}

	// Daily average income (last 30 days)
	err = s.f.pkg.M.FinDash.DB.QueryRow(s.ctx, `
		SELECT COALESCE(AVG(daily_total), 0)
		FROM (
			SELECT DATE(income_date) as day, SUM(income_amount) as daily_total
			FROM income
			WHERE income_date >= CURRENT_DATE - INTERVAL '30 days'
			GROUP BY DATE(income_date)
		) daily_income
	`).Scan(&s.dailyAvgIncome)
	if err != nil {
		return errs.FailedToCalculateIncome
	}

	return nil
}

func (s *service) calculateExpenseMetrics() error {
	// Total expense
	err := s.f.pkg.M.FinDash.DB.QueryRow(s.ctx, `
		SELECT COALESCE(SUM(expense_amount), 0)
		FROM expense
	`).Scan(&s.totalExpense)
	if err != nil {
		return errs.FailedToCalculateExpense
	}

	// Monthly expense (current month)
	err = s.f.pkg.M.FinDash.DB.QueryRow(s.ctx, `
		SELECT COALESCE(SUM(expense_amount), 0)
		FROM expense
		WHERE DATE_TRUNC('month', expense_date) = DATE_TRUNC('month', CURRENT_DATE)
	`).Scan(&s.monthlyExpense)
	if err != nil {
		return errs.FailedToCalculateExpense
	}

	// Daily average expense (last 30 days)
	err = s.f.pkg.M.FinDash.DB.QueryRow(s.ctx, `
		SELECT COALESCE(AVG(daily_total), 0)
		FROM (
			SELECT DATE(expense_date) as day, SUM(expense_amount) as daily_total
			FROM expense
			WHERE expense_date >= CURRENT_DATE - INTERVAL '30 days'
			GROUP BY DATE(expense_date)
		) daily_expense
	`).Scan(&s.dailyAvgExpense)
	if err != nil {
		return errs.FailedToCalculateExpense
	}

	return nil
}

func (s *service) calculateCashFlowMetrics() {
	// Net cash flow
	s.netCashFlow = s.totalIncome - s.totalExpense

	// Savings rate
	if s.totalIncome > 0 {
		s.savingsRate = ((s.totalIncome - s.totalExpense) / s.totalIncome) * 100
	}

	// Stability ratio
	if s.totalExpense > 0 {
		s.stabilityRatio = s.totalIncome / s.totalExpense
	}

	// Stability status
	switch {
	case s.stabilityRatio >= 1.5:
		s.stabilityStatus = "Excellent"
	case s.stabilityRatio >= 1.2:
		s.stabilityStatus = "Good"
	case s.stabilityRatio >= 1.0:
		s.stabilityStatus = "Stable"
	case s.stabilityRatio >= 0.9:
		s.stabilityStatus = "Warning"
	default:
		s.stabilityStatus = "Critical"
	}
}

func (s *service) getTopCategories() error {
	query := `
		SELECT
			expense_type,
			SUM(expense_amount) as total,
			COUNT(*) as count
		FROM expense
		GROUP BY expense_type
		ORDER BY total DESC
		LIMIT 3
	`

	rows, err := s.f.pkg.M.FinDash.DB.Query(s.ctx, query)
	if err != nil {
		return errs.FailedToGetCategories
	}
	defer rows.Close()

	var totalAllCategories float64
	var categories []da.CategorySummary

	// First pass - collect data
	for rows.Next() {
		var cat da.CategorySummary
		var total float64
		err := rows.Scan(&cat.Category, &total, &cat.Count)
		if err != nil {
			return errs.FailedToGetCategories
		}

		// Create Amount object
		cat.Total = []*models.Amount{{
			Amount:         total,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}}

		totalAllCategories += total
		categories = append(categories, cat)
	}

	// Calculate percentages
	for i := range categories {
		if totalAllCategories > 0 {
			categories[i].Percentage = (categories[i].Total[0].Amount / totalAllCategories) * 100
		}
	}

	s.topCategories = categories
	return nil
}

func (s *service) reply() *da.DashboardResponse {
	return &da.DashboardResponse{
		// Income
		TotalIncome: []*models.Amount{{
			Amount:         s.totalIncome,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		MonthlyIncome: []*models.Amount{{
			Amount:         s.monthlyIncome,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		DailyAvgIncome: []*models.Amount{{
			Amount:         s.dailyAvgIncome,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},

		// Expense
		TotalExpense: []*models.Amount{{
			Amount:         s.totalExpense,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		MonthlyExpense: []*models.Amount{{
			Amount:         s.monthlyExpense,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		DailyAvgExpense: []*models.Amount{{
			Amount:         s.dailyAvgExpense,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},

		// Cash Flow
		NetCashFlow: []*models.Amount{{
			Amount:         s.netCashFlow,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		SavingsRate:     s.savingsRate,
		StabilityRatio:  s.stabilityRatio,
		StabilityStatus: s.stabilityStatus,

		// Categories
		TopExpenseCategories: s.topCategories,
	}
}
