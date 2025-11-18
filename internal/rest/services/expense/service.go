package expense

import (
	"github.com/rsmrtk/mybox/internal/rest/services/expense/create"
	"github.com/rsmrtk/mybox/internal/rest/services/expense/delete"
	"github.com/rsmrtk/mybox/internal/rest/services/expense/get"
	"github.com/rsmrtk/mybox/internal/rest/services/expense/list"
	"github.com/rsmrtk/mybox/internal/rest/services/expense/update"
	"github.com/rsmrtk/mybox/pkg"
)

// Service is the expense service facade
type Service struct {
	Get    *get.Facade
	List   *list.Facade
	Create *create.Facade
	Update *update.Facade
	Delete *delete.Facade
}

// New creates a new expense service
func New(f *pkg.Facade) *Service {
	return &Service{
		Get:    get.New(f),
		List:   list.New(f),
		Create: create.New(f),
		Update: update.New(f),
		Delete: delete.New(f),
	}
}
