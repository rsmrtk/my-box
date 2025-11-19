package analytics

import (
	"github.com/rsmrtk/mybox/internal/rest/services/analytics/anomalies"
	"github.com/rsmrtk/mybox/internal/rest/services/analytics/dashboard"
	"github.com/rsmrtk/mybox/internal/rest/services/analytics/growth"
	"github.com/rsmrtk/mybox/internal/rest/services/analytics/topexpenses"
	"github.com/rsmrtk/mybox/internal/rest/services/analytics/trends"
	"github.com/rsmrtk/mybox/pkg"
)

type Service struct {
	Dashboard   *dashboard.Facade
	TopExpenses *topexpenses.Facade
	Growth      *growth.Facade
	Trends      *trends.Facade
	Anomalies   *anomalies.Facade
}

func NewService(pkg *pkg.Facade) *Service {
	return &Service{
		Dashboard:   dashboard.New(pkg),
		TopExpenses: topexpenses.New(pkg),
		Growth:      growth.New(pkg),
		Trends:      trends.New(pkg),
		Anomalies:   anomalies.New(pkg),
	}
}
