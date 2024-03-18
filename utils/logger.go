package utils

import (
	lokihook "github.com/akkuman/logrus-loki-hook"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(appName string) *Logger {
	var log = logrus.New()
	lokiHookConfig := &lokihook.Config{
		URL: "http://localhost:3100/api/prom/push",
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
