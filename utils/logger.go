package utils

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

// Logger return a logger client
func Logger() (logger *logrus.Entry) {
	// logrus.SetLevel(logrus.DebugLevel)

	logrus.SetLevel(logrus.ErrorLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if pc, file, line, ok := runtime.Caller(2); ok {
		f := runtime.FuncForPC(pc)

		return logrus.WithFields(
			logrus.Fields{
				"pc":         pc,
				"file":       file,
				"func.Name":  f.Name(),
				"func.Entry": f.Entry(),
				// "func.Fileline": f.FileLine(pc),
				"line": line})
	}
	return logrus.WithFields(logrus.Fields{})
}
