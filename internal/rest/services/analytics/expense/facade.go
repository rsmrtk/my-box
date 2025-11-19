package expense

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade handles expense analytics operations
type Facade struct {
	ctx context.Context
	req *da.ExpenseAnalyticsRequest
	pkg *pkg.Pkg
	res *da.ExpenseAnalyticsResponse
}

// NewFacade creates a new expense analytics facade
func NewFacade(ctx context.Context, req *da.ExpenseAnalyticsRequest, pkg *pkg.Pkg) *Facade {
	return &Facade{
		ctx: ctx,
		req: req,
		pkg: pkg,
		res: &da.ExpenseAnalyticsResponse{Success: false},
	}
}

// Handle processes the expense analytics request
func (f *Facade) Handle() *da.ExpenseAnalyticsResponse {
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

// service handles the expense analytics logic
type service struct {
	ctx  context.Context
	req  *da.ExpenseAnalyticsRequest
	f    *Facade
	data *da.ExpenseAnalyticsData
}

func (s *service) analyze() error {
	s.data = &da.ExpenseAnalyticsData{}

	switch s.req.AnalysisType {
	case "monthly":
		return s.getMonthlyExpense()
	case "categories":
		return s.getExpenseByCategory()
	case "trends":
		return s.getExpenseTrends()
	case "anomalies":
		return s.getExpenseAnomalies()
	case "share_of_wallet":
		return s.getShareOfWallet()
	default:
		// Get all analytics if no specific type
		s.getMonthlyExpense()
		s.getExpenseByCategory()
		s.getExpenseTrends()
	}

	return nil
}

func (s *service) getMonthlyExpense() error {
	query := `
		SELECT * FROM v_monthly_expense
		WHERE ($1::DATE IS NULL OR month >= $1)
		  AND ($2::DATE IS NULL OR month <= $2)
		ORDER BY month DESC
		LIMIT 12
	`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, s.req.StartDate, s.req.EndDate)
	if err != nil {
		return fmt.Errorf("failed to get monthly expense: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var data da.MonthlyExpenseData
		var expenseTypes sql.NullString

		err := rows.Scan(
			&data.Month,
			&data.TransactionCount,
			&data.TotalAmount,
			&data.AvgAmount,
			&data.MinAmount,
			&data.MaxAmount,
			&expenseTypes,
		)
		if err != nil {
			return fmt.Errorf("failed to scan monthly expense: %w", err)
		}

		// Parse expense types from PostgreSQL array
		if expenseTypes.Valid {
			data.ExpenseTypes = parsePostgresArray(expenseTypes.String)
		}

		s.data.MonthlyExpense = &data
		break // Get only the most recent month
	}

	return nil
}

func (s *service) getExpenseByCategory() error {
	query := `
		SELECT * FROM v_expense_by_category
		WHERE ($1::DATE IS NULL OR 1=1)
		ORDER BY total_amount DESC
	`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, s.req.StartDate)
	if err != nil {
		return fmt.Errorf("failed to get expense by category: %w", err)
	}
	defer rows.Close()

	var categories []da.CategoryExpenseData
	for rows.Next() {
		var data da.CategoryExpenseData
		err := rows.Scan(
			&data.Category,
			&data.TransactionCount,
			&data.TotalAmount,
			&data.AvgAmount,
			&data.MinAmount,
			&data.MaxAmount,
			&data.PercentageOfTotal,
		)
		if err != nil {
			return fmt.Errorf("failed to scan expense category: %w", err)
		}

		categories = append(categories, data)
	}

	s.data.ExpenseByCategory = categories
	return nil
}

func (s *service) getExpenseTrends() error {
	monthsBack := s.req.MonthsBack
	if monthsBack == 0 {
		monthsBack = 6
	}

	query := `SELECT * FROM fn_expense_trend_analysis($1)`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, monthsBack)
	if err != nil {
		return fmt.Errorf("failed to get expense trends: %w", err)
	}
	defer rows.Close()

	var trends []da.ExpenseTrendData
	for rows.Next() {
		var data da.ExpenseTrendData
		var previousExpense, changeAmount, changePercentage, movingAvg sql.NullFloat64

		err := rows.Scan(
			&data.MonthDate,
			&data.TotalExpense,
			&previousExpense,
			&changeAmount,
			&changePercentage,
			&data.TrendDirection,
			&movingAvg,
		)
		if err != nil {
			return fmt.Errorf("failed to scan expense trend: %w", err)
		}

		if previousExpense.Valid {
			data.PreviousMonthExpense = previousExpense.Float64
		}
		if changeAmount.Valid {
			data.ChangeAmount = changeAmount.Float64
		}
		if changePercentage.Valid {
			data.ChangePercentage = changePercentage.Float64
		}
		if movingAvg.Valid {
			data.MovingAvg3Months = movingAvg.Float64
		}

		trends = append(trends, data)
	}

	s.data.ExpenseTrends = trends
	return nil
}

func (s *service) getExpenseAnomalies() error {
	thresholdFactor := s.req.ThresholdFactor
	if thresholdFactor == 0 {
		thresholdFactor = 1.5
	}

	lookbackMonths := 3

	query := `SELECT * FROM fn_expense_anomalies($1, $2)`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, thresholdFactor, lookbackMonths)
	if err != nil {
		return fmt.Errorf("failed to get expense anomalies: %w", err)
	}
	defer rows.Close()

	var anomalies []da.ExpenseAnomalyData
	for rows.Next() {
		var data da.ExpenseAnomalyData
		var expenseID string

		err := rows.Scan(
			&expenseID,
			&data.ExpenseName,
			&data.ExpenseAmount,
			&data.ExpenseType,
			&data.ExpenseDate,
			&data.CategoryAvg,
			&data.DeviationFactor,
			&data.AnomalyScore,
		)
		if err != nil {
			return fmt.Errorf("failed to scan expense anomaly: %w", err)
		}

		// Parse UUID
		data.ExpenseID, _ = uuid.Parse(expenseID)
		anomalies = append(anomalies, data)
	}

	s.data.ExpenseAnomalies = anomalies
	return nil
}

func (s *service) getShareOfWallet() error {
	query := `SELECT * FROM fn_share_of_wallet($1, $2)`

	rows, err := s.f.pkg.M.DB.QueryContext(s.ctx, query, s.req.StartDate, s.req.EndDate)
	if err != nil {
		return fmt.Errorf("failed to get share of wallet: %w", err)
	}
	defer rows.Close()

	var shares []da.ShareOfWalletData
	for rows.Next() {
		var data da.ShareOfWalletData
		err := rows.Scan(
			&data.Category,
			&data.TotalAmount,
			&data.TransactionCount,
			&data.AvgTransaction,
			&data.SharePercentage,
			&data.CumulativePercentage,
			&data.CategoryRank,
		)
		if err != nil {
			return fmt.Errorf("failed to scan share of wallet: %w", err)
		}

		shares = append(shares, data)
	}

	s.data.ShareOfWallet = shares
	return nil
}

func (s *service) reply() *da.ExpenseAnalyticsResponse {
	return &da.ExpenseAnalyticsResponse{
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

// TopExpenseFacade handles top expense requests
type TopExpenseFacade struct {
	ctx context.Context
	req *da.TopExpensesRequest
	pkg *pkg.Pkg
}

// NewTopExpenseFacade creates a new top expense facade
func NewTopExpenseFacade(ctx context.Context, req *da.TopExpensesRequest, pkg *pkg.Pkg) *TopExpenseFacade {
	return &TopExpenseFacade{
		ctx: ctx,
		req: req,
		pkg: pkg,
	}
}

// Handle processes the top expense request
func (f *TopExpenseFacade) Handle() *da.TopExpensesResponse {
	query := `SELECT * FROM fn_top_expenses($1, $2, $3)`

	rows, err := f.pkg.M.DB.QueryContext(f.ctx, query, f.req.Limit, f.req.StartDate, f.req.EndDate)
	if err != nil {
		return &da.TopExpensesResponse{
			Success: false,
			Error:   fmt.Sprintf("failed to get top expenses: %v", err),
		}
	}
	defer rows.Close()

	var items []da.TopExpenseItem
	for rows.Next() {
		var item da.TopExpenseItem
		var expenseID string

		err := rows.Scan(
			&expenseID,
			&item.ExpenseName,
			&item.ExpenseAmount,
			&item.ExpenseType,
			&item.ExpenseDate,
			&item.RankPosition,
		)
		if err != nil {
			return &da.TopExpensesResponse{
				Success: false,
				Error:   fmt.Sprintf("failed to scan top expense: %v", err),
			}
		}

		// Parse UUID
		item.ExpenseID, _ = uuid.Parse(expenseID)
		items = append(items, item)
	}

	return &da.TopExpensesResponse{
		Success: true,
		Data:    items,
	}
}

// TopCategoriesFacade handles top expense categories requests
type TopCategoriesFacade struct {
	ctx context.Context
	req *da.TopExpenseCategoriesRequest
	pkg *pkg.Pkg
}

// NewTopCategoriesFacade creates a new top categories facade
func NewTopCategoriesFacade(ctx context.Context, req *da.TopExpenseCategoriesRequest, pkg *pkg.Pkg) *TopCategoriesFacade {
	return &TopCategoriesFacade{
		ctx: ctx,
		req: req,
		pkg: pkg,
	}
}

// Handle processes the top categories request
func (f *TopCategoriesFacade) Handle() *da.TopExpenseCategoriesResponse {
	query := `SELECT * FROM fn_top_expense_categories($1, $2, $3)`

	rows, err := f.pkg.M.DB.QueryContext(f.ctx, query, f.req.Limit, f.req.StartDate, f.req.EndDate)
	if err != nil {
		return &da.TopExpenseCategoriesResponse{
			Success: false,
			Error:   fmt.Sprintf("failed to get top expense categories: %v", err),
		}
	}
	defer rows.Close()

	var items []da.TopExpenseCategoryItem
	for rows.Next() {
		var item da.TopExpenseCategoryItem
		err := rows.Scan(
			&item.Category,
			&item.TotalAmount,
			&item.TransactionCount,
			&item.AvgAmount,
			&item.PercentageOfTotal,
			&item.RankPosition,
		)
		if err != nil {
			return &da.TopExpenseCategoriesResponse{
				Success: false,
				Error:   fmt.Sprintf("failed to scan top category: %v", err),
			}
		}

		items = append(items, item)
	}

	return &da.TopExpenseCategoriesResponse{
		Success: true,
		Data:    items,
	}
}
