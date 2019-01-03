package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/octowhale/iPicka/util"
)

func main() {

	// logrus.Debugln(config.Storage)
	// Logger := logger.LogConfig{ENV: "debug"}
	// Logger.SetEnv(Logger.ENV)
	// Logger.SetLevel(Logger.Level)

	var HTTPSSchema string
	if config.Storage.HTTPS {
		HTTPSSchema = "https://"
	} else {
		HTTPSSchema = "http://"
	}

	// file := "/data/tmp/naruto.jpg"
	targets := os.Args[1:]

	for _, target := range targets {
		logrus.Debugf("target = %v", target)
		if ok, _ := util.IsDirectory(target); ok {
			logrus.Debugf("%v enter DirMode", target)
			DirMode(target, HTTPSSchema)
		} else {
			Once(target, HTTPSSchema)
		}
	}

}
