package anomalies

import (
	"context"

	da "github.com/rsmrtk/mybox/internal/rest/domain/analytics"
	"github.com/rsmrtk/mybox/pkg"
)

type Facade struct {
	pkg *pkg.Facade
}

func New(pkg *pkg.Facade) *Facade {
	return &Facade{pkg: pkg}
}

func (f *Facade) Handle(ctx context.Context, req *da.AnomaliesRequest) (*da.AnomaliesResponse, error) {
	serv := &service{
		ctx: ctx,
		req: req,
		f:   f,
	}

	if err := serv.detect(); err != nil {
		return nil, err
	}
	return serv.reply(), nil
}
