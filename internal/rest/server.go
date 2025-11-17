package rest

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rsmrtk/mybox/internal/rest/middlewares"
	"github.com/rsmrtk/mybox/internal/rest/services"
	"github.com/rsmrtk/mybox/pkg"
)

type Server struct {
	addr   string
	cert   string
	key    string
	server *gin.Engine
}

type ServerOptions struct {
	Facade   *pkg.Facade
	Services *services.Services
}

type tlsErrorFilter struct {
	facade *pkg.Facade
}

func (w *tlsErrorFilter) Write(p []byte) (n int, err error) {
	logMessage := string(p)

	if strings.Contains(logMessage, "TLS handshake error") && strings.Contains(logMessage, "EOF") {
		if w.facade != nil {
			w.facade.Log.Debug("Filtered TLS handshake EOF error (likely health check or incomplete connection)", nil)
		}
		return len(p), nil
	}

	return os.Stderr.Write(p)
}

func NewServer(o ServerOptions) (*Server, error) {
	engine := gin.New()
	engine.Use(middlewares.CORSMiddleware())
	engine.Use(middlewares.ErrorMiddleware(o.Facade))

	engine.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	engine.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	incomes := engine.Group("/income", middlewares.CORSMiddleware())
	{
		c := controll
	}
}
