package update

import (
	"context"

	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade is the update expense facade
type Facade struct {
	pkg *pkg.Facade
}

// New creates a new update expense facade
func New(pkg *pkg.Facade) *Facade {
	return &Facade{
		pkg: pkg,
	}
}

// Handle handles the update expense request
func (f *Facade) Handle(ctx context.Context, req *expense.UpdateRequest) (*expense.UpdateResponse, error) {
	s := &service{
		ctx: ctx,
		req: req,
		f:   f,
	}

	if err := s.update(); err != nil {
		return nil, err
	}

	return s.reply(), nil
}
