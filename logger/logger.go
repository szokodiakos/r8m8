package logger

import (
	"github.com/evalphobia/logrus_sentry"
	colorable "github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Setup logger
func Setup(sentryDSN string, logFormat string) {
	if logFormat == "json" {
		log.Formatter = &logrus.JSONFormatter{}
	} else {
		log.Formatter = &logrus.TextFormatter{ForceColors: true}
		log.Out = colorable.NewColorableStdout()
	}

	if sentryDSN != "" {
		hook, err := logrus_sentry.NewSentryHook(sentryDSN, []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})

		if err == nil {
			log.Hooks.Add(hook)
			log.Info("Using Sentry Logger")
		}
	}
}

// Get func
func Get() *logrus.Logger {
	return log
}
