package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	er "github.com/rsmrtk/fd-er"
	di "github.com/rsmrtk/mybox/internal/rest/domain/income"
	"github.com/rsmrtk/mybox/internal/rest/services/income"
)

type IncomeController struct {
	service *income.Service
}

func NewEstimateController(service *income.Service) *IncomeController {
	return &IncomeController{service: service}
}

func (c *IncomeController) Get(ctx *gin.Context) {
	var req di.GetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = er.NewHTTPError(http.StatusBadRequest).SetInternal(fmt.Errorf("failed to bind request: %w", err))
		_ = ctx.Error(err)
		return
	}

	res, err := c.service.Get.Handle(ctx, &req)
	if err != nil {
		err = er.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to handle request: %w", err))
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *IncomeController) Create(ctx *gin.Context) {
	var req di.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = er.NewHTTPError(http.StatusBadRequest).SetInternal(fmt.Errorf("failed to bind request: %w", err))
		ctx.Set("failed_request", req)
		_ = ctx.Error(err)
		return
	}

	res, err := c.service.Create.Handle(ctx, &req)
	if err != nil {
		err = er.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to handle request: %w", err))
		ctx.Set("failed_request", req)
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *IncomeController) Update(ctx *gin.Context) {
	var req di.UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = er.NewHTTPError(http.StatusBadRequest).SetInternal(fmt.Errorf("failed to bind request: %w", err))
		ctx.Set("failed_request", req)
		_ = ctx.Error(err)
		return
	}

	res, err := c.service.Update.Handle(ctx, &req)
	if err != nil {
		err = er.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to handle request: %w", err))
		ctx.Set("failed_request", req)
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *IncomeController) List(ctx *gin.Context) {
	var req di.ListRequest
	// For GET requests with query parameters, we can use ShouldBindQuery
	// or just create a default request
	req = di.ListRequest{}

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

	res, err := c.service.List.Handle(ctx, &req)
	if err != nil {
		err = er.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to handle request: %w", err))
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *IncomeController) Delete(ctx *gin.Context) {
	var req di.DeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = er.NewHTTPError(http.StatusBadRequest).SetInternal(fmt.Errorf("failed to bind request: %w", err))
		ctx.Set("failed_request", req)
		_ = ctx.Error(err)
		return
	}

	res, err := c.service.Delete.Handle(ctx, &req)
	if err != nil {
		err = er.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to handle request: %w", err))
		ctx.Set("failed_request", req)
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
