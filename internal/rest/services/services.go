package services

import (
	"github.com/rsmrtk/mybox/internal/rest/services/income"
	"github.com/rsmrtk/mybox/pkg"
)

type Options struct {
	Pkg *pkg.Facade
}

type Services struct {
	Income *income.Service
}

func NewService(opts Options) *Services {
	return &Services{
		Income: income.NewService(opts.Pkg),
	}
}
