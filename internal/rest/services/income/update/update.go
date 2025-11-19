package update

import (
	"context"

	"github.com/rsmrtk/mybox/internal/rest/domain/income"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade is the update income facade
type Facade struct {
	pkg *pkg.Facade
}

// New creates a new update income facade
func New(pkg *pkg.Facade) *Facade {
	return &Facade{
		pkg: pkg,
	}
}

// Handle handles the update income request
func (f *Facade) Handle(ctx context.Context, req *income.UpdateRequest) (*income.UpdateResponse, error) {
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
