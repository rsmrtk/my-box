package analytics

import (
	"time"

	"github.com/google/uuid"
)

// IncomeAnalyticsRequest represents the request for income analytics
type IncomeAnalyticsRequest struct {
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	MonthsBack   int        `json:"months_back,omitempty"`
	AnalysisType string     `json:"analysis_type"` // monthly, annual, growth, statistics, forecast
}

// IncomeAnalyticsResponse represents the response for income analytics
type IncomeAnalyticsResponse struct {
	Success bool                 `json:"success"`
	Data    *IncomeAnalyticsData `json:"data,omitempty"`
	Error   string               `json:"error,omitempty"`
}

// IncomeAnalyticsData contains all income analytics data
type IncomeAnalyticsData struct {
	MonthlyIncome    *MonthlyIncomeData    `json:"monthly_income,omitempty"`
	AnnualIncome     *AnnualIncomeData     `json:"annual_income,omitempty"`
	IncomeGrowth     *IncomeGrowthData     `json:"income_growth,omitempty"`
	IncomeStatistics *IncomeStatisticsData `json:"income_statistics,omitempty"`
	IncomeForecast   *IncomeForecastData   `json:"income_forecast,omitempty"`
}

// MonthlyIncomeData represents monthly income summary
type MonthlyIncomeData struct {
	Month            time.Time `json:"month"`
	TransactionCount int       `json:"transaction_count"`
	TotalAmount      float64   `json:"total_amount"`
	AvgAmount        float64   `json:"avg_amount"`
	MinAmount        float64   `json:"min_amount"`
	MaxAmount        float64   `json:"max_amount"`
	IncomeTypes      []string  `json:"income_types"`
}

// AnnualIncomeData represents annual income summary
type AnnualIncomeData struct {
	Year             int     `json:"year"`
	TransactionCount int     `json:"transaction_count"`
	TotalAmount      float64 `json:"total_amount"`
	AvgAmount        float64 `json:"avg_amount"`
	MinAmount        float64 `json:"min_amount"`
	MaxAmount        float64 `json:"max_amount"`
}

// IncomeGrowthData represents income growth analysis
type IncomeGrowthData struct {
	PeriodDate           time.Time `json:"period_date"`
	PeriodType           string    `json:"period_type"`
	TotalIncome          float64   `json:"total_income"`
	PreviousPeriodIncome float64   `json:"previous_period_income"`
	GrowthAmount         float64   `json:"growth_amount"`
	GrowthPercentage     float64   `json:"growth_percentage"`
}

// IncomeStatisticsData represents income statistics
type IncomeStatisticsData struct {
	AvgIncomeNMonths  float64 `json:"avg_income_n_months"`
	AvgDailyIncome    float64 `json:"avg_daily_income"`
	AvgWeeklyIncome   float64 `json:"avg_weekly_income"`
	AvgMonthlyIncome  float64 `json:"avg_monthly_income"`
	PeriodDescription string  `json:"period_description"`
}

// IncomeForecastData represents income forecast
type IncomeForecastData struct {
	Forecasts []ForecastItem `json:"forecasts"`
}

// ForecastItem represents a single forecast entry
type ForecastItem struct {
	ForecastMonth   time.Time `json:"forecast_month"`
	PredictedIncome float64   `json:"predicted_income"`
	ConfidenceLevel string    `json:"confidence_level"`
	BasedOnMonths   int       `json:"based_on_months"`
}

// TopIncomeRequest represents the request for top income transactions
type TopIncomeRequest struct {
	Limit     int        `json:"limit,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

// TopIncomeResponse represents the response for top income transactions
type TopIncomeResponse struct {
	Success bool            `json:"success"`
	Data    []TopIncomeItem `json:"data,omitempty"`
	Error   string          `json:"error,omitempty"`
}

// TopIncomeItem represents a single top income entry
type TopIncomeItem struct {
	IncomeID     uuid.UUID `json:"income_id"`
	IncomeName   string    `json:"income_name"`
	IncomeAmount float64   `json:"income_amount"`
	IncomeType   string    `json:"income_type"`
	IncomeDate   time.Time `json:"income_date"`
	RankPosition int       `json:"rank_position"`
}
