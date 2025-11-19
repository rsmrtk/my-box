package analytics

import (
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

// ============================================================================
// DASHBOARD
// ============================================================================

type DashboardRequest struct {
	// No parameters needed for dashboard
}

type DashboardResponse struct {
	// Income metrics
	TotalIncome    []*models.Amount `json:"total_income"`
	MonthlyIncome  []*models.Amount `json:"monthly_income"`
	DailyAvgIncome []*models.Amount `json:"daily_avg_income"`

	// Expense metrics
	TotalExpense    []*models.Amount `json:"total_expense"`
	MonthlyExpense  []*models.Amount `json:"monthly_expense"`
	DailyAvgExpense []*models.Amount `json:"daily_avg_expense"`

	// Cash flow metrics
	NetCashFlow     []*models.Amount `json:"net_cash_flow"`
	SavingsRate     float64          `json:"savings_rate"`
	StabilityRatio  float64          `json:"stability_ratio"`
	StabilityStatus string           `json:"stability_status"`

	// Top categories
	TopExpenseCategories []CategorySummary `json:"top_expense_categories"`
}

// ============================================================================
// TOP EXPENSES
// ============================================================================

type TopExpensesRequest struct {
	Limit int `json:"limit"`
}

type TopExpensesResponse struct {
	Categories []CategorySummary `json:"categories"`
}

type CategorySummary struct {
	Category   string           `json:"category"`
	Total      []*models.Amount `json:"total"`
	Count      int              `json:"count"`
	Percentage float64          `json:"percentage"`
}

// ============================================================================
// INCOME GROWTH
// ============================================================================

type GrowthRequest struct {
	// No parameters needed, using current and previous month
}

type GrowthResponse struct {
	CurrentMonth     []*models.Amount `json:"current_month"`
	PreviousMonth    []*models.Amount `json:"previous_month"`
	GrowthAmount     []*models.Amount `json:"growth_amount"`
	GrowthPercentage float64          `json:"growth_percentage"`
}

// ============================================================================
// EXPENSE TRENDS
// ============================================================================

type TrendsRequest struct {
	Months int `json:"months"`
}

type TrendsResponse struct {
	Trends []MonthlyTrend `json:"trends"`
}

type MonthlyTrend struct {
	Month            models.Date      `json:"month"`
	Total            []*models.Amount `json:"total"`
	Count            int              `json:"count"`
	Change           []*models.Amount `json:"change,omitempty"`
	ChangePercentage float64          `json:"change_percentage,omitempty"`
}

// ============================================================================
// EXPENSE ANOMALIES
// ============================================================================

type AnomaliesRequest struct {
	Threshold float64 `json:"threshold"`
}

type AnomaliesResponse struct {
	Anomalies []AnomalyItem `json:"anomalies"`
}

type AnomalyItem struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Amount          []*models.Amount `json:"amount"`
	Type            string           `json:"type"`
	Date            models.Date      `json:"date"`
	CategoryAverage []*models.Amount `json:"category_average"`
	DeviationFactor float64          `json:"deviation_factor"`
	Status          string           `json:"status"` // Critical, High, Medium
}
