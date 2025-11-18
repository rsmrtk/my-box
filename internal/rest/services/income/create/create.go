package create

import (
	"context"

	"github.com/rsmrtk/mybox/internal/rest/domain/income"
	"github.com/rsmrtk/mybox/pkg"
)

type Facade struct {
	pkg *pkg.Facade
}

func New(pkg *pkg.Facade) *Facade {
	return &Facade{pkg: pkg}
}

func (f *Facade) Handle(ctx context.Context, req *income.CreateRequest) (*income.CreateResponse, error) {
	serv := &service{
		ctx: ctx,
		req: req,
		f:   f,
	}

	if err := serv.create(); err != nil {
		return nil, err
	}
	return serv.reply(), nil
}
