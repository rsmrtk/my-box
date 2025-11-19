package analytics

import (
	"time"
)

// CashFlowAnalyticsRequest represents the request for cash flow analytics
type CashFlowAnalyticsRequest struct {
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	ForecastDays int        `json:"forecast_days,omitempty"`
	MonthsBack   int        `json:"months_back,omitempty"`
	AnalysisType string     `json:"analysis_type"` // summary, daily, forecast, stability, emergency_fund
}

// CashFlowAnalyticsResponse represents the response for cash flow analytics
type CashFlowAnalyticsResponse struct {
	Success bool                   `json:"success"`
	Data    *CashFlowAnalyticsData `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}

// CashFlowAnalyticsData contains all cash flow analytics data
type CashFlowAnalyticsData struct {
	CashFlowSummary       []CashFlowSummaryData    `json:"cash_flow_summary,omitempty"`
	DailyCashFlow         []DailyCashFlowData      `json:"daily_cash_flow,omitempty"`
	CashFlowForecast      []CashFlowForecastData   `json:"cash_flow_forecast,omitempty"`
	FinancialStability    []FinancialStabilityData `json:"financial_stability,omitempty"`
	EmergencyFundAnalysis []EmergencyFundData      `json:"emergency_fund_analysis,omitempty"`
}

// CashFlowSummaryData represents cash flow summary
type CashFlowSummaryData struct {
	Month                 time.Time `json:"month"`
	TotalIncome           float64   `json:"total_income"`
	TotalExpense          float64   `json:"total_expense"`
	NetCashFlow           float64   `json:"net_cash_flow"`
	IncomeExpenseRatio    float64   `json:"income_expense_ratio"`
	SavingsRatePercentage float64   `json:"savings_rate_percentage"`
}

// DailyCashFlowData represents daily cash flow
type DailyCashFlowData struct {
	Date    time.Time `json:"date"`
	Income  float64   `json:"income"`
	Expense float64   `json:"expense"`
	NetFlow float64   `json:"net_flow"`
}

// CashFlowForecastData represents cash flow forecast
type CashFlowForecastData struct {
	ForecastDate      time.Time `json:"forecast_date"`
	PredictedIncome   float64   `json:"predicted_income"`
	PredictedExpense  float64   `json:"predicted_expense"`
	PredictedNetFlow  float64   `json:"predicted_net_flow"`
	ConfidenceLevel   string    `json:"confidence_level"`
	CumulativeNetFlow float64   `json:"cumulative_net_flow"`
}

// FinancialStabilityData represents financial stability metrics
type FinancialStabilityData struct {
	MetricName     string  `json:"metric_name"`
	MetricValue    float64 `json:"metric_value"`
	Interpretation string  `json:"interpretation"`
	Status         string  `json:"status"`
}

// EmergencyFundData represents emergency fund analysis
type EmergencyFundData struct {
	MetricName     string  `json:"metric_name"`
	MetricValue    float64 `json:"metric_value"`
	Recommendation string  `json:"recommendation"`
}

// DashboardSummaryRequest represents the request for dashboard summary
type DashboardSummaryRequest struct {
	PeriodStart *time.Time `json:"period_start,omitempty"`
	PeriodEnd   *time.Time `json:"period_end,omitempty"`
}

// DashboardSummaryResponse represents the response for dashboard summary
type DashboardSummaryResponse struct {
	Success bool                   `json:"success"`
	Data    []DashboardSummaryItem `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}

// DashboardSummaryItem represents a single dashboard summary metric
type DashboardSummaryItem struct {
	Section          string  `json:"section"`
	Metric           string  `json:"metric"`
	Value            float64 `json:"value"`
	PercentageChange float64 `json:"percentage_change,omitempty"`
	Trend            string  `json:"trend,omitempty"`
	Description      string  `json:"description"`
}

// FinancialHealthRequest represents the request for financial health check
type FinancialHealthRequest struct {
	StartDate  *time.Time `json:"start_date,omitempty"`
	EndDate    *time.Time `json:"end_date,omitempty"`
	MonthsBack int        `json:"months_back,omitempty"`
}

// FinancialHealthResponse represents the response for financial health check
type FinancialHealthResponse struct {
	Success bool                 `json:"success"`
	Data    *FinancialHealthData `json:"data,omitempty"`
	Error   string               `json:"error,omitempty"`
}

// FinancialHealthData contains financial health metrics
type FinancialHealthData struct {
	OverallScore    float64        `json:"overall_score"`
	HealthStatus    string         `json:"health_status"`
	Metrics         []HealthMetric `json:"metrics"`
	Recommendations []string       `json:"recommendations"`
	Alerts          []HealthAlert  `json:"alerts,omitempty"`
}

// HealthMetric represents a single health metric
type HealthMetric struct {
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	Score       float64 `json:"score"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
}

// HealthAlert represents a financial health alert
type HealthAlert struct {
	Level     string    `json:"level"` // critical, warning, info
	Message   string    `json:"message"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
