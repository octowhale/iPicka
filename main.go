package main

import (
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)
	// logrus.SetLevel(logrus.InfoLevel)
	// logrus.SetLevel(logrus.ErrorLevel)

	logrus.Debugln(config.Storage)

	var HTTPSSchema string
	if config.Storage.HTTPS {
		HTTPSSchema = "https://"
	} else {
		HTTPSSchema = "http://"
	}

	file := "/data/tmp/naruto.jpg"

	Once(file, HTTPSSchema)
}
