package main

import (
	"time"

	lokihook "github.com/akkuman/logrus-loki-hook"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	lokiHookConfig := &lokihook.Config{
		// the loki api url
		URL: "http://localhost:3100/api/prom/push",
		// (optional, default: severity) the label's key to distinguish log's level, it will be added to Labels map
		// LevelName: "severity",
		// the labels which will be sent to loki, contains the {levelname: level}
		Labels: map[string]string{
			"application": "broker",
		},
		// BatchWait: time.Nanosecond,
		BatchEntriesNumber: 1,
	}
	hook, err := lokihook.NewHook(lokiHookConfig)
	if err != nil {
		log.Error(err)
	} else {
		log.AddHook(hook)
	}
}

func main() {
	for i := 0; i < 20000; i++ {
		log.Infof("romio %d times", i)
		time.Sleep(time.Millisecond * 500)
	}
}
