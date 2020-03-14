package rdbms

import (
	"context"
	"database/sql"
	"github.com/maestre3d/placehub-payment/internal/shared/infrastructure/config"
	"github.com/sony/gobreaker"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresPool(lc fx.Lifecycle, config *config.Config, logger *zap.SugaredLogger) *sql.DB {

	// Circuit Breaker avoids failure leaks
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:          "rdbms.postgres",
		MaxRequests:   0,
		Interval:      0,
		Timeout:       0,
		ReadyToTrip:   nil,
		OnStateChange: nil,
	})
	logger.Infow("Postgres_Circuit_Breaker",
		"status", "started",
	)

	var dbPool *sql.DB
	// Open Postgres connection
	dbGeneric, err := cb.Execute(func() (i interface{}, err error) {
		db, err := sql.Open("postgres", config.RDBMSConnection)
		if err != nil {
			return nil, err
		}

		if err = db.Ping(); err != nil {
			return nil, err
		}

		db.SetConnMaxLifetime(8 * time.Second)
		db.SetMaxOpenConns(4)
		db.SetMaxIdleConns(1)

		logger.Infow("Postgres_Pool",
			"status", "started",
		)

		return db, nil
	})
	if err != nil {
		panic(err)
	}

	dbPool = dbGeneric.(*sql.DB)

	lc.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			logger.Infow("Postgres_Pool",
				"status", "closed",
			)
			return dbPool.Close()
		},
	})

	return dbPool
}

func NewPostgresConn(db *sql.DB, logger *zap.SugaredLogger) *sql.Conn {
	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}

	logger.Infow("Postgres_Connection",
		"status", "started",
	)

	return conn
}
