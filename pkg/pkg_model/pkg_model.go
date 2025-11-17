package pkg_model

import (
	"context"

	"github.com/rsmrtk/smartlg/logger"

	dbModelFinDash "github.com/rsmrtk/db-fd-model"
)

type Models struct {
	FinDash *dbModelFinDash.Model
}

func New(ctx context.Context, spannerUrlFinDash string, lg *logger.Logger) (*Models, error) {
	findashInstance, err := dbModelFinDash.New(ctx, &dbModelFinDash.Options{
		SpannerUrl: spannerUrlFinDash,
		Log:        lg,
	})
	if err != nil {
		return nil, err
	}

	return &Models{
		FinDash: findashInstance,
	}, nil
}
