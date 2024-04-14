package logger

import (
	"fmt"

	lokihook "github.com/akkuman/logrus-loki-hook"
	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/sirupsen/logrus"
)

var (
	LOKI_HOST = env.GetEnv("GRAFANA_LOKI_HOST", "localhost")
	LOKI_PORT = env.GetEnv("GRAFANA_LOKI_PORT", "3100")
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(appName string) *Logger {
	var log = logrus.New()
	lokiHookConfig := &lokihook.Config{
		URL: fmt.Sprintf("http://%s:%s/api/prom/push", LOKI_HOST, LOKI_PORT),
		Labels: map[string]string{
			"application": appName,
		},
		BatchEntriesNumber: 1,
	}
	hook, err := lokihook.NewHook(lokiHookConfig)
	if err != nil {
		log.Error(err)
	} else {
		log.AddHook(hook)
	}
	return &Logger{log}
}
