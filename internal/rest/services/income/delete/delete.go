package delete

import (
	"context"

	"github.com/rsmrtk/mybox/internal/rest/domain/income"
	"github.com/rsmrtk/mybox/pkg"
)

// Facade is the delete income facade
type Facade struct {
	pkg *pkg.Facade
}

// New creates a new delete income facade
func New(pkg *pkg.Facade) *Facade {
	return &Facade{
		pkg: pkg,
	}
}

// Handle handles the delete income request
func (f *Facade) Handle(ctx context.Context, req *income.DeleteRequest) (*income.DeleteResponse, error) {
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
