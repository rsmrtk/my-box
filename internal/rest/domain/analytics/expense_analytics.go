package analytics

import (
	"time"

	"github.com/google/uuid"
)

// ExpenseAnalyticsRequest represents the request for expense analytics
type ExpenseAnalyticsRequest struct {
	StartDate       *time.Time `json:"start_date,omitempty"`
	EndDate         *time.Time `json:"end_date,omitempty"`
	MonthsBack      int        `json:"months_back,omitempty"`
	AnalysisType    string     `json:"analysis_type"`              // monthly, categories, trends, anomalies, share_of_wallet
	ThresholdFactor float64    `json:"threshold_factor,omitempty"` // For anomaly detection
}

// ExpenseAnalyticsResponse represents the response for expense analytics
type ExpenseAnalyticsResponse struct {
	Success bool                  `json:"success"`
	Data    *ExpenseAnalyticsData `json:"data,omitempty"`
	Error   string                `json:"error,omitempty"`
}

// ExpenseAnalyticsData contains all expense analytics data
type ExpenseAnalyticsData struct {
	MonthlyExpense    *MonthlyExpenseData   `json:"monthly_expense,omitempty"`
	ExpenseByCategory []CategoryExpenseData `json:"expense_by_category,omitempty"`
	ExpenseTrends     []ExpenseTrendData    `json:"expense_trends,omitempty"`
	ExpenseAnomalies  []ExpenseAnomalyData  `json:"expense_anomalies,omitempty"`
	ShareOfWallet     []ShareOfWalletData   `json:"share_of_wallet,omitempty"`
}

// MonthlyExpenseData represents monthly expense summary
type MonthlyExpenseData struct {
	Month            time.Time `json:"month"`
	TransactionCount int       `json:"transaction_count"`
	TotalAmount      float64   `json:"total_amount"`
	AvgAmount        float64   `json:"avg_amount"`
	MinAmount        float64   `json:"min_amount"`
	MaxAmount        float64   `json:"max_amount"`
	ExpenseTypes     []string  `json:"expense_types"`
}

// CategoryExpenseData represents expense data by category
type CategoryExpenseData struct {
	Category          string  `json:"category"`
	TransactionCount  int     `json:"transaction_count"`
	TotalAmount       float64 `json:"total_amount"`
	AvgAmount         float64 `json:"avg_amount"`
	MinAmount         float64 `json:"min_amount"`
	MaxAmount         float64 `json:"max_amount"`
	PercentageOfTotal float64 `json:"percentage_of_total"`
}

// ExpenseTrendData represents expense trend analysis
type ExpenseTrendData struct {
	MonthDate            time.Time `json:"month_date"`
	TotalExpense         float64   `json:"total_expense"`
	PreviousMonthExpense float64   `json:"previous_month_expense"`
	ChangeAmount         float64   `json:"change_amount"`
	ChangePercentage     float64   `json:"change_percentage"`
	TrendDirection       string    `json:"trend_direction"`
	MovingAvg3Months     float64   `json:"moving_avg_3_months"`
}

// ExpenseAnomalyData represents anomaly detection results
type ExpenseAnomalyData struct {
	ExpenseID       uuid.UUID `json:"expense_id"`
	ExpenseName     string    `json:"expense_name"`
	ExpenseAmount   float64   `json:"expense_amount"`
	ExpenseType     string    `json:"expense_type"`
	ExpenseDate     time.Time `json:"expense_date"`
	CategoryAvg     float64   `json:"category_avg"`
	DeviationFactor float64   `json:"deviation_factor"`
	AnomalyScore    string    `json:"anomaly_score"`
}

// ShareOfWalletData represents share of wallet analysis
type ShareOfWalletData struct {
	Category             string  `json:"category"`
	TotalAmount          float64 `json:"total_amount"`
	TransactionCount     int     `json:"transaction_count"`
	AvgTransaction       float64 `json:"avg_transaction"`
	SharePercentage      float64 `json:"share_percentage"`
	CumulativePercentage float64 `json:"cumulative_percentage"`
	CategoryRank         int     `json:"category_rank"`
}

// TopExpensesRequest represents the request for top expenses
type TopExpensesRequest struct {
	Limit     int        `json:"limit,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

// TopExpensesResponse represents the response for top expenses
type TopExpensesResponse struct {
	Success bool             `json:"success"`
	Data    []TopExpenseItem `json:"data,omitempty"`
	Error   string           `json:"error,omitempty"`
}

// TopExpenseItem represents a single top expense entry
type TopExpenseItem struct {
	ExpenseID     uuid.UUID `json:"expense_id"`
	ExpenseName   string    `json:"expense_name"`
	ExpenseAmount float64   `json:"expense_amount"`
	ExpenseType   string    `json:"expense_type"`
	ExpenseDate   time.Time `json:"expense_date"`
	RankPosition  int       `json:"rank_position"`
}

// TopExpenseCategoriesRequest represents the request for top expense categories
type TopExpenseCategoriesRequest struct {
	Limit     int        `json:"limit,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

// TopExpenseCategoriesResponse represents the response for top expense categories
type TopExpenseCategoriesResponse struct {
	Success bool                     `json:"success"`
	Data    []TopExpenseCategoryItem `json:"data,omitempty"`
	Error   string                   `json:"error,omitempty"`
}

// TopExpenseCategoryItem represents a single top expense category
type TopExpenseCategoryItem struct {
	Category          string  `json:"category"`
	TotalAmount       float64 `json:"total_amount"`
	TransactionCount  int     `json:"transaction_count"`
	AvgAmount         float64 `json:"avg_amount"`
	PercentageOfTotal float64 `json:"percentage_of_total"`
	RankPosition      int     `json:"rank_position"`
}
