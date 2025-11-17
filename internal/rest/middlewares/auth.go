package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	er "github.com/rsmrtk/fd-er"
	"github.com/rsmrtk/mybox/pkg/utils"
)

const headerAPIKey = "X-API-Key"

var apiKeys = map[string]*string{
	"Q9mTf4BxeJ7pK2vZsR1uML8aG0cYhWdXn3oCbPVtF6qUwEiHr5jSkDgNOPYBvL": utils.ValueToPtr("e89462b1-cf26-4dc4-8ee4-ffa9194443e8"),
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(headerAPIKey)
		customerID, exists := apiKeys[key]
		if !exists {
			_ = c.Error(er.NewHTTPError(http.StatusUnauthorized, "invalid API key"))
			c.Abort()
			return
		}
		if customerID != nil {
			utils.GinAuthSetCtx(c, *customerID)
		}
		c.Next()
	}
}
