package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	er "github.com/rsmrtk/fd-er"
	"github.com/rsmrtk/mybox/internal/rest/services/income"
)

type IncomeController struct {
	service *income.Service
}

func NewEstimateController(service *income.Service) *IncomeController {
	return &IncomeController{service: service}
}

func (c *IncomeController) Get(ctx *gin.Context) {
	var req domainestimate.GetRequest
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

func (c *IncomeController) GetToll(ctx *gin.Context) {
	var req domainestimate.GetTollRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = er.NewHTTPError(http.StatusBadRequest).SetInternal(fmt.Errorf("failed to bind request: %w", err))
		ctx.Set("failed_request", req)
		_ = ctx.Error(err)
		return
	}

	res, err := c.service.GetToll.Handle(ctx, &req)
	if err != nil {
		err = er.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to handle request: %w", err))
		ctx.Set("failed_request", req)
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
