package pkg_model

import (
	"context"

	dbModelFinDash "github.com/rsmrtk/db-fd-model"
	"github.com/rsmrtk/smartlg/logger"
)

type Models struct {
	FinDash *dbModelFinDash.Model
}

func New(ctx context.Context, postgresURL string, lg *logger.Logger) (*Models, error) {
	finDashInstance, err := dbModelFinDash.New(ctx, &dbModelFinDash.Options{
		PostgresURL: postgresURL,
		Log:         lg,
	})
	if err != nil {
		return nil, err
	}

	return &Models{
		FinDash: finDashInstance,
	}, nil
}
