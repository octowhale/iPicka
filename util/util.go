package util

import (
	"crypto/md5"
	"fmt"
	"io"
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

func GetMd5(file string) (string, error) {

	if ok, err := IsFileExist(file); !ok {
		logrus.Errorln(err)
		return "", err
	}

	// b, err := ioutil.ReadFile(file)
	fi, err := os.Open(file)
	if err != nil {
		logrus.Errorln(err)
		return "", err
	}

	h := md5.New()
	io.Copy(h, fi)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func GetEnvDefault(key, defaults string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return defaults
	}
	return v
}
