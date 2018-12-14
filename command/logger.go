package command

import (
	"github.com/sirupsen/logrus"
)

// Logger return a logger client
func Logger() (logger *logrus.Entry) {
	// logrus.SetLevel(logrus.DebugLevel)

	logrus.SetLevel(logrus.ErrorLevel)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	return logrus.WithFields(logrus.Fields{})
}
