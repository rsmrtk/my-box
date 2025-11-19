package cashflow

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade handles cash flow analytics operations
type Facade struct {
	ctx context.Context
	req *da.CashFlowAnalyticsRequest
	pkg *pkg.Facade
	res *da.CashFlowAnalyticsResponse
}

// NewFacade creates a new cash flow analytics facade
func NewFacade(ctx context.Context, req *da.CashFlowAnalyticsRequest, pkg *pkg.Facade) *Facade {
	return &Facade{
		ctx: ctx,
		req: req,
		pkg: pkg,
		res: &da.CashFlowAnalyticsResponse{Success: false},
	}
}

// Handle processes the cash flow analytics request
func (f *Facade) Handle() *da.CashFlowAnalyticsResponse {
	s := &service{
		ctx: f.ctx,
		req: f.req,
		f:   f,
	}

	if err := s.analyze(); err != nil {
		f.res.Error = err.Error()
		return f.res
	}

	return s.reply()
}

// service handles the cash flow analytics logic
type service struct {
	ctx  context.Context
	req  *da.CashFlowAnalyticsRequest
	f    *Facade
	data *da.CashFlowAnalyticsData
}

func (s *service) analyze() error {
	s.data = &da.CashFlowAnalyticsData{}

	switch s.req.AnalysisType {
	case "summary":
		return s.getCashFlowSummary()
	case "daily":
		return s.getDailyCashFlow()
	case "stability":
		return s.getFinancialStability()
	default:
		// Get all analytics if no specific type
		s.getCashFlowSummary()
		s.getFinancialStability()
	}

	return nil
}

func (s *service) getCashFlowSummary() error {
	query := `
		SELECT * FROM v_cash_flow_summary
		WHERE ($1::DATE IS NULL OR month >= $1)
		  AND ($2::DATE IS NULL OR month <= $2)
		ORDER BY month DESC
		LIMIT 12
	`

	rows, err := s.f.pkg.M.FinDash.DB.QueryContext(s.ctx, query, s.req.StartDate, s.req.EndDate)
	if err != nil {
		return fmt.Errorf("failed to get cash flow summary: %w", err)
	}
	defer rows.Close()

	var summaries []da.CashFlowSummaryData
	for rows.Next() {
		var data da.CashFlowSummaryData
		var incomeExpenseRatio, savingsRate sql.NullFloat64

		err := rows.Scan(
			&data.Month,
			&data.TotalIncome,
			&data.TotalExpense,
			&data.NetCashFlow,
			&incomeExpenseRatio,
			&savingsRate,
		)
		if err != nil {
			return fmt.Errorf("failed to scan cash flow summary: %w", err)
		}

		if incomeExpenseRatio.Valid {
			data.IncomeExpenseRatio = incomeExpenseRatio.Float64
		}
		if savingsRate.Valid {
			data.SavingsRatePercentage = savingsRate.Float64
		}

		summaries = append(summaries, data)
	}

	s.data.CashFlowSummary = summaries
	return nil
}

func (s *service) getDailyCashFlow() error {
	query := `
		SELECT * FROM v_daily_cash_flow
		WHERE ($1::DATE IS NULL OR date >= $1)
		  AND ($2::DATE IS NULL OR date <= $2)
		ORDER BY date DESC
		LIMIT 90
	`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, s.req.StartDate, s.req.EndDate)
	if err != nil {
		return fmt.Errorf("failed to get daily cash flow: %w", err)
	}
	defer rows.Close()

	var dailyFlows []da.DailyCashFlowData
	for rows.Next() {
		var data da.DailyCashFlowData
		err := rows.Scan(
			&data.Date,
			&data.Income,
			&data.Expense,
			&data.NetFlow,
		)
		if err != nil {
			return fmt.Errorf("failed to scan daily cash flow: %w", err)
		}

		dailyFlows = append(dailyFlows, data)
	}

	s.data.DailyCashFlow = dailyFlows
	return nil
}

func (s *service) getFinancialStability() error {
	monthsBack := s.req.MonthsBack
	if monthsBack == 0 {
		monthsBack = 6
	}

	query := `SELECT * FROM fn_financial_stability_coefficient($1)`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, monthsBack)
	if err != nil {
		return fmt.Errorf("failed to get financial stability: %w", err)
	}
	defer rows.Close()

	var stability []da.FinancialStabilityData
	for rows.Next() {
		var data da.FinancialStabilityData
		err := rows.Scan(
			&data.MetricName,
			&data.MetricValue,
			&data.Interpretation,
			&data.Status,
		)
		if err != nil {
			return fmt.Errorf("failed to scan financial stability: %w", err)
		}

		stability = append(stability, data)
	}

	s.data.FinancialStability = stability
	return nil
}

func (s *service) reply() *da.CashFlowAnalyticsResponse {
	return &da.CashFlowAnalyticsResponse{
		Success: true,
		Data:    s.data,
	}
}

// DashboardSummaryFacade handles dashboard summary requests
type DashboardSummaryFacade struct {
	ctx context.Context
	req *da.DashboardSummaryRequest
	pkg *pkg.Pkg
}

// NewDashboardSummaryFacade creates a new dashboard summary facade
func NewDashboardSummaryFacade(ctx context.Context, req *da.DashboardSummaryRequest, pkg *pkg.Pkg) *DashboardSummaryFacade {
	return &DashboardSummaryFacade{
		ctx: ctx,
		req: req,
		pkg: pkg,
	}
}

// Handle processes the dashboard summary request
func (f *DashboardSummaryFacade) Handle() *da.DashboardSummaryResponse {
	query := `SELECT * FROM fn_financial_dashboard_summary($1, $2)`

	rows, err := f.pkg.M.DB.QueryContext(f.ctx, query, f.req.PeriodStart, f.req.PeriodEnd)
	if err != nil {
		return &da.DashboardSummaryResponse{
			Success: false,
			Error:   fmt.Sprintf("failed to get dashboard summary: %v", err),
		}
	}
	defer rows.Close()

	var items []da.DashboardSummaryItem
	for rows.Next() {
		var item da.DashboardSummaryItem
		var percentageChange, value sql.NullFloat64
		var trend sql.NullString

		err := rows.Scan(
			&item.Section,
			&item.Metric,
			&value,
			&percentageChange,
			&trend,
			&item.Description,
		)
		if err != nil {
			return &da.DashboardSummaryResponse{
				Success: false,
				Error:   fmt.Sprintf("failed to scan dashboard summary: %v", err),
			}
		}

		if value.Valid {
			item.Value = value.Float64
		}
		if percentageChange.Valid {
			item.PercentageChange = percentageChange.Float64
		}
		if trend.Valid {
			item.Trend = trend.String
		}

		items = append(items, item)
	}

	return &da.DashboardSummaryResponse{
		Success: true,
		Data:    items,
	}
}
