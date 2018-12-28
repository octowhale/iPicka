package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

func IsFileExist(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func IsDirectory(path string) (bool, error) {
	fi, err := os.Stat(path)

	if err != nil {
		logrus.Debugln(path, err)
		return false, err
	}

	if fi.IsDir() {
		logrus.Debugln(path, "is a directory")
		return true, nil
	}

	logrus.Debugln(path, "is not a directory")
	return false, nil
}
