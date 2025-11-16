package app

import "github.com/rsmrtk/mybox/pkg"

type App struct {
	pkg *pkg.Facade

	servicesGRPC *grpcservices.Service
	servicesREST *restservices.Service
}
