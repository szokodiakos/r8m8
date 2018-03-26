package logger

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Setup logger
func Setup(sentryDSN string) {
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
