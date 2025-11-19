package services

import (
	"github.com/rsmrtk/mybox/internal/rest/services/analytics"
	"github.com/rsmrtk/mybox/internal/rest/services/expense"
	"github.com/rsmrtk/mybox/internal/rest/services/income"
	"github.com/rsmrtk/mybox/pkg"
)

type Options struct {
	Pkg *pkg.Facade
}

type Services struct {
	Income    *income.Service
	Expense   *expense.Service
	Analytics *analytics.Service
}

func NewService(opts Options) *Services {
	return &Services{
		Income:    income.NewService(opts.Pkg),
		Expense:   expense.New(opts.Pkg),
		Analytics: analytics.NewService(opts.Pkg),
	}
}
