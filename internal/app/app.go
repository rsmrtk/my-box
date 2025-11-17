package app

import (
	"context"
	"os/signal"
	"syscall"

	restservices "github.com/rsmrtk/mybox/internal/rest/services"
	"github.com/rsmrtk/mybox/pkg"
)

type App struct {
	pkg *pkg.Facade

	//servicesGRPC *grpcservices.Service
	servicesREST *restservices.Services
}

func Run() {
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer cancelFunc()

	app := new(App)

	pkg, err := pkg.New(ctx)
	if err != nil {
		panic(err)
	}

	app.pkg = pkg
	//app.servicesGRPC = grpcservices.NewService(grpcservices.Options{Pkg: pkg})
	app.servicesREST = restservices.NewService(restservices.Options{Pkg: pkg})

	pkg.Log.Infof("Starting REST server...")

}
