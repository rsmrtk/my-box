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
		ac := controllers.NewAnalyticsController(o.Facade)

		// Income Analytics
		incomeAnalytics := analytics.Group("/income")
		{
			incomeAnalytics.GET("", ac.GetIncomeAnalytics)         // General income analytics
			incomeAnalytics.GET("/top", ac.GetTopIncomes)          // Top incomes
			incomeAnalytics.GET("/growth", ac.GetIncomeGrowth)     // Income growth analysis
			incomeAnalytics.GET("/forecast", ac.GetIncomeForecast) // Income forecast
		}

		// Expense Analytics
		expenseAnalytics := analytics.Group("/expense")
		{
			expenseAnalytics.GET("", ac.GetExpenseAnalytics)                    // General expense analytics
			expenseAnalytics.GET("/top", ac.GetTopExpenses)                     // Top expenses
			expenseAnalytics.GET("/top-categories", ac.GetTopExpenseCategories) // Top expense categories
			expenseAnalytics.GET("/anomalies", ac.GetExpenseAnomalies)          // Expense anomaly detection
			expenseAnalytics.GET("/trends", ac.GetExpenseTrends)                // Expense trends
			expenseAnalytics.GET("/share-of-wallet", ac.GetShareOfWallet)       // Share of wallet analysis
		}

		// Cash Flow Analytics
		cashFlowAnalytics := analytics.Group("/cashflow")
		{
			cashFlowAnalytics.GET("/summary", ac.GetCashFlowSummary)              // Cash flow summary
			cashFlowAnalytics.GET("/forecast", ac.GetCashFlowForecast)            // Cash flow forecast
			cashFlowAnalytics.GET("/stability", ac.GetFinancialStability)         // Financial stability
			cashFlowAnalytics.GET("/emergency-fund", ac.GetEmergencyFundAnalysis) // Emergency fund analysis
		}

		// Dashboard & Health
		analytics.GET("/dashboard", ac.GetDashboardSummary)       // Comprehensive dashboard summary
		analytics.GET("/financial-health", ac.GetFinancialHealth) // Financial health check
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
