package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/internal/rest/services/analytics"
)

// AnalyticsController handles analytics endpoints
type AnalyticsController struct {
	service *analytics.Service
}

// NewAnalyticsController creates a new analytics controller
func NewAnalyticsController(s *analytics.Service) *AnalyticsController {
	return &AnalyticsController{
		service: s,
	}
}

// GetDashboard handles dashboard request
// GET /analytics/dashboard
func (c *AnalyticsController) GetDashboard(ctx *gin.Context) {
	req := &da.DashboardRequest{}

	resp, err := c.service.Dashboard.Handle(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// GetTopExpenses handles top expenses request
// GET /analytics/expenses/top?limit=3
func (c *AnalyticsController) GetTopExpenses(ctx *gin.Context) {
	limit := 3 // default value

	req := &da.TopExpensesRequest{
		Limit: limit,
	}

	resp, err := c.service.TopExpenses.Handle(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// GetIncomeGrowth handles income growth request
// GET /analytics/income/growth
func (c *AnalyticsController) GetIncomeGrowth(ctx *gin.Context) {
	req := &da.GrowthRequest{}

	resp, err := c.service.Growth.Handle(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// GetExpenseTrends handles expense trends request
// GET /analytics/expenses/trends?months=6
func (c *AnalyticsController) GetExpenseTrends(ctx *gin.Context) {
	months := 6 // default value

	req := &da.TrendsRequest{
		Months: months,
	}

	resp, err := c.service.Trends.Handle(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// GetAnomalies handles expense anomalies request
// GET /analytics/expenses/anomalies?threshold=1.5
func (c *AnalyticsController) GetAnomalies(ctx *gin.Context) {
	var threshold float64 = 1.5 // default

	req := &da.AnomaliesRequest{
		Threshold: threshold,
	}

	resp, err := c.service.Anomalies.Handle(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}
