package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/maestre3d/placehub-payment/internal/shared/infrastructure/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func NewMux(lc fx.Lifecycle, logger *zap.SugaredLogger, config *config.Config) *gin.Engine {
	mux := gin.Default()
	// mux := http.NewServeMux()
	server := &http.Server{
		Addr:    config.HTTPPort,
		Handler: mux,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Infow("Starting HTTP Server",
					"context", "server.NewMux",
				)
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		},
	)

	return mux
}
