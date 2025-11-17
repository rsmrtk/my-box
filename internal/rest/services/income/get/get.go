package get

import (
	"context"

	"github.com/rsmrtk/mybox/pkg"
)

type Facade struct {
	pkg *pkg.Facade
}

func New(pkg *pkg.Facade) *Facade {
	return &Facade{pkg: pkg}
}

func (f *Facade) Handler(ctx context.Context, req *income.GetRequest) error {
	serv := &service{
		ctx: ctx,
		req: req,
		f:   f,
	}

	if err := serv.find(); err != nil {
		return nil, err
	}

}
