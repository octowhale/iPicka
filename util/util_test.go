package util

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_util(t *testing.T) {

	// logrus.SetLevel(logrus.DebugLevel)

	path := "/data/tmp/naruto.jp"

	b, err := IsDirectory(path)
	if err != nil {
		logrus.Errorln("ERRRRR:", err)
	}
	logrus.Infoln(b)
}
