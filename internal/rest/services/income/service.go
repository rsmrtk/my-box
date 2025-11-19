package income

import (
	"github.com/rsmrtk/mybox/internal/rest/services/income/create"
	"github.com/rsmrtk/mybox/internal/rest/services/income/delete"
	"github.com/rsmrtk/mybox/internal/rest/services/income/get"
	"github.com/rsmrtk/mybox/internal/rest/services/income/list"
	"github.com/rsmrtk/mybox/internal/rest/services/income/update"
	"github.com/rsmrtk/mybox/pkg"
)

type Service struct {
	Get    *get.Facade
	List   *list.Facade
	Create *create.Facade
	Update *update.Facade
	Delete *delete.Facade
}

func NewService(pkg *pkg.Facade) *Service {
	return &Service{
		Get:    get.New(pkg),
		List:   list.New(pkg),
		Create: create.New(pkg),
		Update: update.New(pkg),
		Delete: delete.New(pkg),
	}
}
