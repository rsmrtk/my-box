package list

import (
	"context"

	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade is the list expense facade
type Facade struct {
	pkg *pkg.Facade
}

// New creates a new list expense facade
func New(pkg *pkg.Facade) *Facade {
	return &Facade{
		pkg: pkg,
	}
}

// Handle handles the list expense request
func (f *Facade) Handle(ctx context.Context, req *expense.ListRequest) (*expense.ListResponse, error) {
	s := &service{
		ctx: ctx,
		req: req,
		f:   f,
	}

	if err := s.list(); err != nil {
		return nil, err
	}

	return s.reply(), nil
}
