package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/internal/rest/services/analytics"
	"github.com/rsmrtk/smartlg/logger"
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
// GET /analytics/dashboard?currency=USD
func (c *AnalyticsController) GetDashboard(ctx *gin.Context) {
	// Parse currency parameter
	currency := ctx.DefaultQuery("currency", "USD")

	req := &da.DashboardRequest{
		Currency: currency,
	}

	resp, err := c.service.Dashboard.Handle(ctx.Request.Context(), req)
	if err != nil {
		logger.Error("Failed to get dashboard data", map[string]any{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve dashboard data. Please try again later",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// GetTopExpenses handles top expenses request
// GET /analytics/expenses/top?limit=3&currency=USD
func (c *AnalyticsController) GetTopExpenses(ctx *gin.Context) {
	// Parse limit parameter with validation
	limit := 3 // default value
	if limitStr := ctx.Query("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			logger.Error("Failed to parse limit parameter", map[string]any{
				"error": err.Error(),
				"value": limitStr,
			})
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid limit parameter. Must be a positive integer",
			})
			return
		}

		// Validate limit range
		if parsedLimit <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Limit must be greater than 0",
			})
			return
		}
		if parsedLimit > 100 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Limit cannot exceed 100",
			})
			return
		}
		limit = parsedLimit
	}

	// Parse currency parameter
	currency := ctx.DefaultQuery("currency", "USD")

	req := &da.TopExpensesRequest{
		Limit:    limit,
		Currency: currency,
	}

	resp, err := c.service.TopExpenses.Handle(ctx.Request.Context(), req)
	if err != nil {
		logger.Error("Failed to get top expenses", map[string]any{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve top expenses. Please try again later",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// GetIncomeGrowth handles income growth request
// GET /analytics/income/growth?currency=USD
func (c *AnalyticsController) GetIncomeGrowth(ctx *gin.Context) {
	// Parse currency parameter
	currency := ctx.DefaultQuery("currency", "USD")

	req := &da.GrowthRequest{
		Currency: currency,
	}

	resp, err := c.service.Growth.Handle(ctx.Request.Context(), req)
	if err != nil {
		logger.Error("Failed to get income growth", map[string]any{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve income growth data. Please try again later",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// GetExpenseTrends handles expense trends request
// GET /analytics/expenses/trends?months=6&currency=USD
func (c *AnalyticsController) GetExpenseTrends(ctx *gin.Context) {
	// Parse months parameter with validation
	months := 6 // default value
	if monthsStr := ctx.Query("months"); monthsStr != "" {
		parsedMonths, err := strconv.Atoi(monthsStr)
		if err != nil {
			logger.Error("Failed to parse months parameter", map[string]any{
				"error": err.Error(),
				"value": monthsStr,
			})
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid months parameter. Must be a positive integer",
			})
			return
		}

		// Validate months range
		if parsedMonths <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Months must be greater than 0",
			})
			return
		}
		if parsedMonths > 24 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Months cannot exceed 24",
			})
			return
		}
		months = parsedMonths
	}

	// Parse currency parameter
	currency := ctx.DefaultQuery("currency", "USD")

	req := &da.TrendsRequest{
		Months:   months,
		Currency: currency,
	}

	resp, err := c.service.Trends.Handle(ctx.Request.Context(), req)
	if err != nil {
		logger.Error("Failed to get expense trends", map[string]any{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve expense trends. Please try again later",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// GetAnomalies handles expense anomalies request
// GET /analytics/expenses/anomalies?threshold=1.5&currency=USD
func (c *AnalyticsController) GetAnomalies(ctx *gin.Context) {
	// Parse threshold parameter with validation
	threshold := 1.5 // default value
	if thresholdStr := ctx.Query("threshold"); thresholdStr != "" {
		parsedThreshold, err := strconv.ParseFloat(thresholdStr, 64)
		if err != nil {
			logger.Error("Failed to parse threshold parameter", map[string]any{
				"error": err.Error(),
				"value": thresholdStr,
			})
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid threshold parameter. Must be a positive number",
			})
			return
		}

		// Validate threshold range
		if parsedThreshold <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Threshold must be greater than 0",
			})
			return
		}
		if parsedThreshold > 10 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Threshold cannot exceed 10",
			})
			return
		}
		threshold = parsedThreshold
	}

	// Parse currency parameter
	currency := ctx.DefaultQuery("currency", "USD")

	req := &da.AnomaliesRequest{
		Threshold: threshold,
		Currency:  currency,
	}

	resp, err := c.service.Anomalies.Handle(ctx.Request.Context(), req)
	if err != nil {
		logger.Error("Failed to get anomalies", map[string]any{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to detect expense anomalies. Please try again later",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}
