package bootstrap

import (
	"github.com/maestre3d/placehub-payment/internal/payment/application"
	rdbms2 "github.com/maestre3d/placehub-payment/internal/payment/infrastructure/persistence/rdbms"
	server "github.com/maestre3d/placehub-payment/internal/presentation/delivery/http"
	"github.com/maestre3d/placehub-payment/internal/presentation/delivery/http/handler"
	"github.com/maestre3d/placehub-payment/internal/shared/infrastructure/config"
	"github.com/maestre3d/placehub-payment/internal/shared/infrastructure/logger"
	"github.com/maestre3d/placehub-payment/internal/shared/infrastructure/persistence/rdbms"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			logger.NewLogger,
			logger.NewZapLogger,
			logger.NewSugarLogger,
			config.NewConfig,
			rdbms.NewPostgresPool,
			rdbms.NewPostgresConn,
			server.NewMux,
			server.NewProxyRouter,
			rdbms2.NewPaymentRepository,
			application.NewPaymentUseCase,
		),
		fx.Invoke(handler.InitPaymentHandler),
	)
}
