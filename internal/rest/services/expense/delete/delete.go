package delete

import (
	"context"

	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade is the delete expense facade
type Facade struct {
	pkg *pkg.Facade
}

// New creates a new delete expense facade
func New(pkg *pkg.Facade) *Facade {
	return &Facade{
		pkg: pkg,
	}
}

// Handle handles the delete expense request
func (f *Facade) Handle(ctx context.Context, req *expense.DeleteRequest) (*expense.DeleteResponse, error) {
	s := &service{
		ctx: ctx,
		req: req,
		f:   f,
	}

	if err := s.delete(); err != nil {
		return nil, err
	}

	return s.reply(), nil
}
