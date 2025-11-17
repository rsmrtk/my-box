package income

import (
	"github.com/rsmrtk/mybox/internal/rest/services/income/create"
	"github.com/rsmrtk/mybox/internal/rest/services/income/get"
	"github.com/rsmrtk/mybox/pkg"
)

type Service struct {
	Get    *get.Facade
	Create *create.Facade
}

func NewService(pkg *pkg.Facade) *Service {
	return &Service{
		Get:    get.New(pkg),
		Create: create.New(pkg),
	}
}
