package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
	expenseService "github.com/rsmrtk/mybox/internal/rest/services/expense"
)

// ExpenseController handles expense-related HTTP requests
type ExpenseController struct {
	service *expenseService.Service
}

// NewExpenseController creates a new expense controller
func NewExpenseController(service *expenseService.Service) *ExpenseController {
	return &ExpenseController{
		service: service,
	}
}

// Get handles GET request for fetching an expense
func (c *ExpenseController) Get(ctx *gin.Context) {
	var req expense.GetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := c.service.Get.Handle(ctx.Request.Context(), &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// Create handles POST request for creating an expense
func (c *ExpenseController) Create(ctx *gin.Context) {
	var req expense.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := c.service.Create.Handle(ctx.Request.Context(), &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// Update handles PUT request for updating an expense
func (c *ExpenseController) Update(ctx *gin.Context) {
	var req expense.UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := c.service.Update.Handle(ctx.Request.Context(), &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// List handles GET request for listing all expenses
func (c *ExpenseController) List(ctx *gin.Context) {
	req := expense.ListRequest{}

	// Parse query parameters
	if limit := ctx.Query("limit"); limit != "" {
		var limitVal int
		fmt.Sscanf(limit, "%d", &limitVal)
		req.Limit = limitVal
	}
	if offset := ctx.Query("offset"); offset != "" {
		var offsetVal int
		fmt.Sscanf(offset, "%d", &offsetVal)
		req.Offset = offsetVal
	}
	if sortBy := ctx.Query("sort_by"); sortBy != "" {
		req.SortBy = sortBy
	}
	if order := ctx.Query("order"); order != "" {
		req.Order = order
	}

	resp, err := c.service.List.Handle(ctx.Request.Context(), &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// Delete handles DELETE request for deleting an expense
func (c *ExpenseController) Delete(ctx *gin.Context) {
	var req expense.DeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := c.service.Delete.Handle(ctx.Request.Context(), &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
