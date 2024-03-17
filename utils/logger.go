package utils

import (
	"context"
	"time"

	zaploki "github.com/paul-milne/zap-loki"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(appName string) *Logger {
	logger, _ := initLogger("http://localhost:3100/api/prom/push", appName)
	defer logger.Sync()
	return &Logger{Logger: logger}
}

func initLogger(lokiAddress string, appName string) (*zap.Logger, error) {
	zapConfig := zap.NewProductionConfig()
	loki := zaploki.New(context.Background(), zaploki.Config{
		Url:          lokiAddress,
		BatchMaxSize: 1000,
		BatchMaxWait: 10 * time.Second,
		Labels:       map[string]string{"app": appName},
	})

	return loki.WithCreateLogger(zapConfig)
}
