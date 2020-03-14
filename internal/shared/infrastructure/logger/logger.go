package logger

import (
	"go.uber.org/zap"
	"log"
	"os"
)

func NewLogger() *log.Logger {
	customLog := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	customLog.Print("Executing NewLogger.")
	return customLog
}

func NewZapLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		panic(err)
	}

	return logger
}

func NewSugarLogger(logger *zap.Logger) *zap.SugaredLogger {
	return logger.Sugar()
}
