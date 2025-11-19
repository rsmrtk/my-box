package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rsmrtk/mybox/internal/rest/controllers"
	"github.com/rsmrtk/mybox/internal/rest/middlewares"
	"github.com/rsmrtk/mybox/internal/rest/services"
	"github.com/rsmrtk/mybox/pkg"
	log "github.com/rsmrtk/smartlg"
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
		c := controllers.NewEstimateController(o.Services.Income)
		incomes.GET("/list", c.List) // List all incomes
		incomes.GET("", c.Get)       // Get single income
		incomes.POST("", c.Create)
		incomes.PUT("", c.Update)
		incomes.DELETE("", c.Delete)
	}

	expenses := engine.Group("/expense", middlewares.CORSMiddleware())
	{
		c := controllers.NewExpenseController(o.Services.Expense)
		expenses.GET("/list", c.List) // List all expenses
		expenses.GET("", c.Get)       // Get single expense
		expenses.POST("", c.Create)
		expenses.PUT("", c.Update)
		expenses.DELETE("", c.Delete)
	}

	// Analytics endpoints
	analytics := engine.Group("/analytics", middlewares.CORSMiddleware())
	{
		ac := controllers.NewAnalyticsController(o.Services.Analytics)

		// Dashboard endpoint
		analytics.GET("/dashboard", ac.GetDashboard) // Dashboard summary with income/expense metrics

		// Expense analytics
		analytics.GET("/expenses/top", ac.GetTopExpenses)      // Top expense categories
		analytics.GET("/expenses/trends", ac.GetExpenseTrends) // Expense trends over time
		analytics.GET("/expenses/anomalies", ac.GetAnomalies)  // Expense anomaly detection

		// Income analytics
		analytics.GET("/income/growth", ac.GetIncomeGrowth) // Income growth analysis
	}

	return &Server{addr: ":9595", cert: o.Facade.Config.TLSCertFile, key: o.Facade.Config.TLSKeyFile, server: engine}, nil
}

func (s *Server) Serve() error {
	errorFilter := &tlsErrorFilter{facade: nil}

	httpSrv := &http.Server{
		Addr:     s.addr,
		Handler:  s.server,
		ErrorLog: log.New(errorFilter, "", log.LstdFlags),
	}

	// Use HTTP for local development when certificates are not provided
	if s.cert == "" || s.key == "" {
		if err := httpSrv.ListenAndServe(); err != nil {
			return fmt.Errorf("internal error: %w", err)
		}
	} else {
		if err := httpSrv.ListenAndServeTLS(s.cert, s.key); err != nil {
			return fmt.Errorf("internal error: %w", err)
		}
	}
	return nil
}

func (s *Server) Shutdown(_ context.Context) error {
	return nil
}
