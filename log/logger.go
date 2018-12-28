package log

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func init() {

	level2 := map[string]logrus.Level{
		"env":    logrus.DebugLevel,
		"online": logrus.ErrorLevel,
	}

	logrus.SetLevel(level2[strings.ToLower(os.Getenv("ENV"))])

}

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
