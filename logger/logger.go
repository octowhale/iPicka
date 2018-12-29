package logger

import (
	"strings"

	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	Level string
	ENV   string
	// log *logrus.SetLevel()
}

func (log *LogConfig) SetLevel(lvl string) {

	level := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
	}

	logrus.SetLevel(level[lvl])
}

func (log *LogConfig) SetEnv(env string) {

	// log
	filenameHook := filename.NewHook()
	filenameHook.Field = "file"
	logrus.AddHook(filenameHook)

	if strings.ToLower(env) == "online" {
		// formatter
		logrus.SetFormatter(&logrus.JSONFormatter{
			// PrettyPrint: true,
		})
	}

	if strings.ToLower(env) == "dev" {
		logrus.SetFormatter(&logrus.TextFormatter{
			// ForceColors: false,
			DisableColors:    true,
			DisableTimestamp: false,
			DisableSorting:   true,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "message"},
		})
	}

}
