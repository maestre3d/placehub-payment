package main

import (
	server "github.com/maestre3d/placehub-payment/internals/presentation/delivery/http"
	"github.com/maestre3d/placehub-payment/internals/presentation/delivery/http/handler"
	"go.uber.org/fx"
	"log"
	"os"
)

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			handler.NewUserHandler,
			server.NewMux,
		),
		fx.Invoke(server.Register),
	)

	app.Run()
}

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	return logger
}
