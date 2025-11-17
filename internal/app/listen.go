package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	lg "github.com/rsmrtk/smartlg/logger"

	"github.com/rsmrtk/mybox/internal/rest"
)

func (app *App) Listen() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	//serverGRPC, err := grpc.NewServer(grpc.ServerOptions{
	//	Facade:   app.pkg,
	//	Services: app.servicesGRPC,
	//})
	//if err != nil {
	//	app.pkg.Log.Fatal("gRPC server error", log.H{"error": err})
	//}

	serverREST, err := rest.NewServer(rest.ServerOptions{
		Facade:   app.pkg,
		Services: app.servicesREST,
	})
	if err != nil {
		app.pkg.Log.Fatal("REST server error", lg.H{"error": err})
	}

	//go func() {
	//	defer cancel()
	//	app.pkg.Log.Infof("gRPC server started")
	//	if err := serverGRPC.Serve(); err != nil {
	//		app.pkg.Log.Fatal("gRPC server error", log.H{"error": err})
	//	}
	//}()
	go func() {
		defer cancel()
		app.pkg.Log.Infof("REST server started")
		if err := serverREST.Serve(); err != nil {
			app.pkg.Log.Fatal("REST server error", lg.H{"error": err})
		}
	}()
	//
	//<-ctx.Done() // Server is stopped.
	//app.pkg.Log.Infof("gRPC server stopped")
	//if err := serverGRPC.Shutdown(context.Background()); err != nil {
	//	app.pkg.Log.Fatal("gRPC server shutdown error", map[string]any{"error": err})
	//}

	app.pkg.Log.Infof("REST server stopped")
	if err := serverREST.Shutdown(context.Background()); err != nil {
		app.pkg.Log.Fatal("REST server shutdown error", map[string]any{"error": err})
	}
}
