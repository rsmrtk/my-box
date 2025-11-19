package get

import (
	"context"

	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade is the get expense facade
type Facade struct {
	pkg *pkg.Facade
}

// New creates a new get expense facade
func New(pkg *pkg.Facade) *Facade {
	return &Facade{
		pkg: pkg,
	}
}

// Handle handles the get expense request
func (f *Facade) Handle(ctx context.Context, req *expense.GetRequest) (*expense.GetResponse, error) {
	s := &service{
		ctx: ctx,
		req: req,
		f:   f,
	}

	if err := s.find(); err != nil {
		return nil, err
	}

	return s.reply(), nil
}
