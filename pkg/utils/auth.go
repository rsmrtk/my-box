package utils

import (
	"context"

	"github.com/gin-gonic/gin"
)

type authCustomerID struct{}

func AuthCtx(ctx context.Context) string {
	return ctx.Value(authCustomerID{}).(string)
}

func AuthSetCtx(ctx context.Context, customerID string) context.Context {
	return context.WithValue(ctx, authCustomerID{}, customerID)
}

const ginAuthCustomerID = "customerID"

func GinAuthSetCtx(ctx *gin.Context, customerID string) {
	ctx.Set(ginAuthCustomerID, customerID)
}

func GinAuthCtx(ctx context.Context) string {
	return ctx.Value(ginAuthCustomerID).(string)
}
