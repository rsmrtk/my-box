package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	er "github.com/rsmrtk/fd-er"
	"github.com/rsmrtk/mybox/pkg"
	lg "github.com/rsmrtk/smartlg/logger"
)

func ErrorMiddleware(pkg *pkg.Facade) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request.
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			var herr *er.HTTPError

			// Get failed request if it exists
			logData := lg.H{}
			if failedReq, exists := c.Get("failed_request"); exists {
				logData["request"] = failedReq
			}
			logData["path"] = c.Request.URL.Path
			logData["method"] = c.Request.Method

			switch {
			case errors.As(err.Err, &herr):
				m := gin.H{"code": herr.Code, "message": herr.Message}
				if herr.Internal != nil {
					m["internal"] = herr.Internal.Error()
				}
				c.AbortWithStatusJSON(herr.Code, gin.H{"error": m})
				logData["error"] = m
				pkg.Log.Error("HTTP error", logData)
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": gin.H{
				"code":     http.StatusInternalServerError,
				"message":  http.StatusText(http.StatusInternalServerError),
				"internal": err.Err.Error(),
			}})
			logData["error"] = err.Err.Error()
			pkg.Log.Error("HTTP error", logData)
		}
	}
}
