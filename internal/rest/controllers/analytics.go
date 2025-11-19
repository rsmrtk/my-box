package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/internal/rest/services/analytics/cashflow"
	"github.com/rsmrtk/mybox/internal/rest/services/analytics/expense"
	"github.com/rsmrtk/mybox/internal/rest/services/analytics/income"
	"github.com/rsmrtk/mybox/pkg"
)

// AnalyticsController handles all analytics endpoints
type AnalyticsController struct {
	pkg *pkg.Facade
}

// NewAnalyticsController creates a new analytics controller
func NewAnalyticsController(p *pkg.Facade) *AnalyticsController {
	return &AnalyticsController{pkg: p}
}

// =============================================================================
// INCOME ANALYTICS ENDPOINTS
// =============================================================================

// GetIncomeAnalytics handles income analytics request
func (ac *AnalyticsController) GetIncomeAnalytics(c *gin.Context) {
	var req da.IncomeAnalyticsRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.IncomeAnalyticsResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Call income analytics service
	facade := income.NewFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// GetTopIncomes handles top incomes request
func (ac *AnalyticsController) GetTopIncomes(c *gin.Context) {
	var req da.TopIncomeRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.TopIncomeResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Default limit if not provided
	if req.Limit == 0 {
		req.Limit = 3
	}

	// Call income analytics service
	facade := income.NewTopIncomeFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// GetIncomeGrowth handles income growth analysis request
func (ac *AnalyticsController) GetIncomeGrowth(c *gin.Context) {
	var req da.IncomeAnalyticsRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.IncomeAnalyticsResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	req.AnalysisType = "growth"

	// Call income analytics service
	facade := income.NewFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// =============================================================================
// EXPENSE ANALYTICS ENDPOINTS
// =============================================================================

// GetExpenseAnalytics handles expense analytics request
func (ac *AnalyticsController) GetExpenseAnalytics(c *gin.Context) {
	var req da.ExpenseAnalyticsRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.ExpenseAnalyticsResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Call expense analytics service
	facade := expense.NewFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// GetTopExpenses handles top expenses request
func (ac *AnalyticsController) GetTopExpenses(c *gin.Context) {
	var req da.TopExpensesRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.TopExpensesResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Default limit if not provided
	if req.Limit == 0 {
		req.Limit = 3
	}

	// Call expense analytics service
	facade := expense.NewTopExpenseFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// GetTopExpenseCategories handles top expense categories request
func (ac *AnalyticsController) GetTopExpenseCategories(c *gin.Context) {
	var req da.TopExpenseCategoriesRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.TopExpenseCategoriesResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Default limit if not provided
	if req.Limit == 0 {
		req.Limit = 3
	}

	// Call expense analytics service
	facade := expense.NewTopCategoriesFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// GetExpenseAnomalies handles expense anomalies detection request
func (ac *AnalyticsController) GetExpenseAnomalies(c *gin.Context) {
	var req da.ExpenseAnalyticsRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.ExpenseAnalyticsResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	req.AnalysisType = "anomalies"

	// Default threshold factor if not provided
	if req.ThresholdFactor == 0 {
		req.ThresholdFactor = 1.5
	}

	// Call expense analytics service
	facade := expense.NewFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// GetExpenseTrends handles expense trends analysis request
func (ac *AnalyticsController) GetExpenseTrends(c *gin.Context) {
	var req da.ExpenseAnalyticsRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.ExpenseAnalyticsResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	req.AnalysisType = "trends"

	// Call expense analytics service
	facade := expense.NewFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// GetShareOfWallet handles share of wallet analysis request
func (ac *AnalyticsController) GetShareOfWallet(c *gin.Context) {
	var req da.ExpenseAnalyticsRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.ExpenseAnalyticsResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	req.AnalysisType = "share_of_wallet"

	// Call expense analytics service
	facade := expense.NewFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// =============================================================================
// CASH FLOW ANALYTICS ENDPOINTS
// =============================================================================

// GetCashFlowSummary handles cash flow summary request
func (ac *AnalyticsController) GetCashFlowSummary(c *gin.Context) {
	var req da.CashFlowAnalyticsRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.CashFlowAnalyticsResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	req.AnalysisType = "summary"

	// Call cash flow analytics service
	facade := cashflow.NewFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// GetFinancialStability handles financial stability analysis request
func (ac *AnalyticsController) GetFinancialStability(c *gin.Context) {
	var req da.CashFlowAnalyticsRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.CashFlowAnalyticsResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	req.AnalysisType = "stability"

	// Default months back if not provided
	if req.MonthsBack == 0 {
		req.MonthsBack = 6
	}

	// Call cash flow analytics service
	facade := cashflow.NewFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}

// GetDashboardSummary handles comprehensive dashboard summary request
func (ac *AnalyticsController) GetDashboardSummary(c *gin.Context) {
	var req da.DashboardSummaryRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, da.DashboardSummaryResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Call cash flow analytics service for dashboard summary
	facade := cashflow.NewDashboardSummaryFacade(c.Request.Context(), &req, ac.pkg)
	response := facade.Handle()

	c.JSON(http.StatusOK, response)
}
