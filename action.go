package main

import (
	"path"
	"strings"

	"github.com/octowhale/iPicka/backend"
	"github.com/octowhale/iPicka/storage"
	"github.com/octowhale/iPicka/util"
	"github.com/sirupsen/logrus"
)

func Once(file string, HTTPSSchema string) {

	fileUrl, err := Upload(file)
	if err != nil {
		// logrus.Errorln(err)
		panic(err)
	}

	logrus.Infof("fileUrl: %s", fileUrl)
	SetDB(HTTPSSchema+fileUrl, file)

}

func Upload(file string) (string, error) {

	storageClient, _ := storage.New(config.Storage)

	if ok, err := util.IsFileExist(file); !ok {
		return "", err
	}

	var object string
	if len(config.Storage.Prefix) != 0 {
		object = strings.Trim(config.Storage.Prefix, "/") + "/" + path.Base(file)
	} else {
		object = path.Base(file)
	}

	var fileUrl string
	if len(config.Storage.CustomDomain) == 0 {
		fileUrl, err := storageClient.Put(object, file)
		if err != nil {
			return "", err
		}
		return fileUrl, nil
	}

	fileUrl = config.Storage.CustomDomain + "/" + object
	logrus.Infof("UPload: %s", fileUrl)

	return fileUrl, nil

}

func SetDB(fileUrl, file string) (string, error) {

	md5sum, err := util.GetMd5(file)
	if err != nil {
		return "", err
	}

	backendClient, _ := backend.New(config.Backend)

	url, err := backendClient.Get(md5sum)
	if err != nil {
		return "", err
	} else {
		// return url

		logrus.Debugln(url)
		return url, nil
	}

	_, err = backendClient.Set(md5sum, fileUrl)
	if err != nil {
		return "", err
	}
	url, _ = backendClient.Get(md5sum)
	logrus.Debugln(url)

	return url, nil
}
