package income

import (
	"context"
	"database/sql"
	"fmt"

	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade handles income analytics operations
type Facade struct {
	ctx context.Context
	req *da.IncomeAnalyticsRequest
	pkg *pkg.Facade
	res *da.IncomeAnalyticsResponse
}

// NewFacade creates a new income analytics facade
func NewFacade(ctx context.Context, req *da.IncomeAnalyticsRequest, pkg *pkg.Facade) *Facade {
	return &Facade{
		ctx: ctx,
		req: req,
		pkg: pkg,
		res: &da.IncomeAnalyticsResponse{Success: false},
	}
}

// Handle processes the income analytics request
func (f *Facade) Handle() *da.IncomeAnalyticsResponse {
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

// service handles the income analytics logic
type service struct {
	ctx  context.Context
	req  *da.IncomeAnalyticsRequest
	f    *Facade
	data *da.IncomeAnalyticsData
}

func (s *service) analyze() error {
	s.data = &da.IncomeAnalyticsData{}

	switch s.req.AnalysisType {
	case "monthly":
		return s.getMonthlyIncome()
	case "annual":
		return s.getAnnualIncome()
	case "growth":
		return s.getIncomeGrowth()
	case "statistics":
		return s.getIncomeStatistics()
	default:
		// Get all analytics if no specific type
		s.getMonthlyIncome()
		s.getIncomeGrowth()
		s.getIncomeStatistics()
	}

	return nil
}

func (s *service) getMonthlyIncome() error {
	query := `
		SELECT * FROM v_monthly_income
		WHERE ($1::DATE IS NULL OR month >= $1)
		  AND ($2::DATE IS NULL OR month <= $2)
		ORDER BY month DESC
		LIMIT 12
	`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, s.req.StartDate, s.req.EndDate)
	if err != nil {
		return fmt.Errorf("failed to get monthly income: %w", err)
	}
	defer rows.Close()

	var monthlyData []da.MonthlyIncomeData
	for rows.Next() {
		var data da.MonthlyIncomeData
		var incomeTypes sql.NullString

		err := rows.Scan(
			&data.Month,
			&data.TransactionCount,
			&data.TotalAmount,
			&data.AvgAmount,
			&data.MinAmount,
			&data.MaxAmount,
			&incomeTypes,
		)
		if err != nil {
			return fmt.Errorf("failed to scan monthly income: %w", err)
		}

		// Parse income types from PostgreSQL array
		if incomeTypes.Valid {
			// Parse PostgreSQL array format: {type1,type2}
			// This is simplified - in production you might want a proper parser
			data.IncomeTypes = parsePostgresArray(incomeTypes.String)
		}

		monthlyData = append(monthlyData, data)
	}

	if len(monthlyData) > 0 {
		s.data.MonthlyIncome = &monthlyData[0]
	}

	return nil
}

func (s *service) getAnnualIncome() error {
	query := `
		SELECT * FROM v_annual_income
		ORDER BY year DESC
		LIMIT 5
	`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query)
	if err != nil {
		return fmt.Errorf("failed to get annual income: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var data da.AnnualIncomeData
		err := rows.Scan(
			&data.Year,
			&data.TransactionCount,
			&data.TotalAmount,
			&data.AvgAmount,
			&data.MinAmount,
			&data.MaxAmount,
		)
		if err != nil {
			return fmt.Errorf("failed to scan annual income: %w", err)
		}

		s.data.AnnualIncome = &data
		break // Get only the most recent year
	}

	return nil
}

func (s *service) getIncomeGrowth() error {
	query := `
		SELECT * FROM fn_income_growth_analysis($1, $2)
		ORDER BY period_date DESC
		LIMIT 12
	`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, s.req.StartDate, s.req.EndDate)
	if err != nil {
		return fmt.Errorf("failed to get income growth: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var data da.IncomeGrowthData
		var previousIncome, growthAmount, growthPercentage sql.NullFloat64

		err := rows.Scan(
			&data.PeriodDate,
			&data.PeriodType,
			&data.TotalIncome,
			&previousIncome,
			&growthAmount,
			&growthPercentage,
		)
		if err != nil {
			return fmt.Errorf("failed to scan income growth: %w", err)
		}

		if previousIncome.Valid {
			data.PreviousPeriodIncome = previousIncome.Float64
		}
		if growthAmount.Valid {
			data.GrowthAmount = growthAmount.Float64
		}
		if growthPercentage.Valid {
			data.GrowthPercentage = growthPercentage.Float64
		}

		s.data.IncomeGrowth = &data
		break // Get only the most recent growth data
	}

	return nil
}

func (s *service) getIncomeStatistics() error {
	monthsBack := s.req.MonthsBack
	if monthsBack == 0 {
		monthsBack = 6
	}

	query := `SELECT * FROM fn_income_statistics($1)`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, monthsBack)
	if err != nil {
		return fmt.Errorf("failed to get income statistics: %w", err)
	}
	defer rows.Close()

	stats := &da.IncomeStatisticsData{}
	for rows.Next() {
		var metricName string
		var metricValue float64
		var periodDesc string

		err := rows.Scan(&metricName, &metricValue, &periodDesc)
		if err != nil {
			return fmt.Errorf("failed to scan income statistics: %w", err)
		}

		switch metricName {
		case fmt.Sprintf("avg_income_%d_months", monthsBack):
			stats.AvgIncomeNMonths = metricValue
		case "avg_daily_income":
			stats.AvgDailyIncome = metricValue
		case "avg_weekly_income":
			stats.AvgWeeklyIncome = metricValue
		case "avg_monthly_income":
			stats.AvgMonthlyIncome = metricValue
		}
		stats.PeriodDescription = periodDesc
	}

	s.data.IncomeStatistics = stats
	return nil
}

func (s *service) reply() *da.IncomeAnalyticsResponse {
	return &da.IncomeAnalyticsResponse{
		Success: true,
		Data:    s.data,
	}
}

// Helper function to parse PostgreSQL array string
func parsePostgresArray(arrayStr string) []string {
	// Remove curly braces
	if len(arrayStr) < 2 {
		return []string{}
	}
	arrayStr = arrayStr[1 : len(arrayStr)-1]

	// Split by comma
	if arrayStr == "" {
		return []string{}
	}

	// Simple split - in production you might need more sophisticated parsing
	// to handle quoted strings with commas
	result := []string{}
	current := ""
	inQuotes := false

	for _, char := range arrayStr {
		if char == '"' {
			inQuotes = !inQuotes
		} else if char == ',' && !inQuotes {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}

// TopIncomeFacade handles top income requests
type TopIncomeFacade struct {
	ctx context.Context
	req *da.TopIncomeRequest
	pkg *pkg.Pkg
}

// NewTopIncomeFacade creates a new top income facade
func NewTopIncomeFacade(ctx context.Context, req *da.TopIncomeRequest, pkg *pkg.Pkg) *TopIncomeFacade {
	return &TopIncomeFacade{
		ctx: ctx,
		req: req,
		pkg: pkg,
	}
}

// Handle processes the top income request
func (f *TopIncomeFacade) Handle() *da.TopIncomeResponse {
	query := `SELECT * FROM fn_top_incomes($1, $2, $3)`

	rows, err := f.pkg.M.DB.QueryContext(f.ctx, query, f.req.Limit, f.req.StartDate, f.req.EndDate)
	if err != nil {
		return &da.TopIncomeResponse{
			Success: false,
			Error:   fmt.Sprintf("failed to get top incomes: %v", err),
		}
	}
	defer rows.Close()

	var items []da.TopIncomeItem
	for rows.Next() {
		var item da.TopIncomeItem
		err := rows.Scan(
			&item.IncomeID,
			&item.IncomeName,
			&item.IncomeAmount,
			&item.IncomeType,
			&item.IncomeDate,
			&item.RankPosition,
		)
		if err != nil {
			return &da.TopIncomeResponse{
				Success: false,
				Error:   fmt.Sprintf("failed to scan top income: %v", err),
			}
		}
		items = append(items, item)
	}

	return &da.TopIncomeResponse{
		Success: true,
		Data:    items,
	}
}
