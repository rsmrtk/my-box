package pkg

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rsmrtk/fd-cfg/config"
	"github.com/rsmrtk/fd-cfg/config/env"
	"github.com/rsmrtk/fd-storage/storage"
	"github.com/rsmrtk/mybox/pkg/jwt"
	"github.com/rsmrtk/mybox/pkg/pkg_model"
	lg "github.com/rsmrtk/smartlg/logger"
)

type Facade struct {
	Log     *lg.Logger
	M       *pkg_model.Models
	Config  *Config
	JWT     jwt.JWT
	Storage *storage.Client
}

type Facades struct {
	ENV     env.ENV
	Log     *lg.Logger
	M       *pkg_model.Models
	Storage *storage.Client
	JWT     jwt.JWT
}

func New(ctx context.Context) (*Facade, error) {
	cfgInstance, err := initCfg(ctx)
	if err != nil {
		return nil, err
	}

	logInstance, err := initLogger(cfgInstance.ENV)
	if err != nil {
		return nil, err
	}

	modelsInstance, err := initModels(ctx, cfgInstance, logInstance)
	if err != nil {
		return nil, err
	}

	bucketInstance, err := initBucket()
	if err != nil {
		return nil, err
	}

	jwtInstance, err := initJWT(cfgInstance)
	if err != nil {
		return nil, err
	}

	facade := &Facade{
		Log:     logInstance,
		M:       modelsInstance,
		Config:  cfgInstance,
		JWT:     jwtInstance,
		Storage: bucketInstance,
	}

	return facade, nil
}

func initLogger(env env.ENV) (*lg.Logger, error) {
	level := lg.DebugLevel
	if env.IsProd {
		level = lg.InfoLevel
	}
	w := io.MultiWriter(os.Stdout)
	l := lg.New(w).With().Timestamp().Logger().Level(level)
	return l, nil
}

func initModels(ctx context.Context, cnfInstance *Config, logInstance *lg.Logger) (*pkg_model.Models, error) {
	modelsInstance, err := pkg_model.New(ctx, cnfInstance.PostgresURL, logInstance)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize models: %w", err)
	}

	return modelsInstance, nil
}

func initJWT(cnfInstance *Config) (jwt.JWT, error) {
	duration, err := time.ParseDuration(cnfInstance.JWTDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWTDuration: %w", err)
	}
	return jwt.New(cnfInstance.JWTSecret, duration), nil
}

func initBucket() (*storage.Client, error) {
	bucketInstance, err := storage.New(&storage.Params{
		BucketNames: storage.BucketNames{
			Public:  "jurny_rider",
			Private: "jurny_rider_public",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bucket: %w", err)
	}
	return bucketInstance, nil
}

func initCfg(ctx context.Context) (*Config, error) {
	env := env.New(os.Getenv("ENV"))

	valueMap, err := config.Load(ctx, &config.Options{
		ENV:           &env,
		SecretConfigs: []config.SecretConfigValue{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Get PostgreSQL connection string from environment
	postgresURL := os.Getenv("POSTGRES_DSN")
	if postgresURL == "" {
		postgresURL = "postgres://postgres:password@localhost:5432/mybox?sslmode=disable"
	}

	c := &Config{
		ENV:         env,
		PostgresURL: postgresURL,
		JWTSecret:   os.Getenv("JWT_SECRET"),
		JWTDuration: os.Getenv("JWT_DURATION"),
		TLSCertFile: valueMap.Get(config.TLSCertFile.Name),
		TLSKeyFile:  valueMap.Get(config.TLSKeyFile.Name),
	}

	return c, nil
}
