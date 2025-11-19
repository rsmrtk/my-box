package create

import (
	"context"

	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade is the create expense facade
type Facade struct {
	pkg *pkg.Facade
}

// New creates a new create expense facade
func New(pkg *pkg.Facade) *Facade {
	return &Facade{
		pkg: pkg,
	}
}

// Handle handles the create expense request
func (f *Facade) Handle(ctx context.Context, req *expense.CreateRequest) (*expense.CreateResponse, error) {
	s := &service{
		ctx: ctx,
		req: req,
		f:   f,
	}

	if err := s.create(); err != nil {
		return nil, err
	}

	return s.reply(), nil
}
