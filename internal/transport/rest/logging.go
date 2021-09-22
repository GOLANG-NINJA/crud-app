package rest

import "github.com/sirupsen/logrus"

func logFields(handler string) logrus.Fields {
	return logrus.Fields{
		"handler": handler,
	}
}

func logError(handler string, err error) {
	logrus.WithFields(logFields(handler)).Error(err)
}
