package bootstrap

import (
	server "github.com/maestre3d/placehub-payment/internals/presentation/delivery/http"
	"github.com/maestre3d/placehub-payment/internals/presentation/delivery/http/handler"
	"github.com/maestre3d/placehub-payment/internals/shared/infrastructure/config"
	"github.com/maestre3d/placehub-payment/internals/shared/infrastructure/logger"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			logger.NewLogger,
			logger.NewZapLogger,
			logger.NewSugarLogger,
			config.NewConfig,
			server.NewMux,
			server.NewProxyRouter,
			handler.NewUserHandler,
		),
		fx.Invoke(server.Register),
	)
}
