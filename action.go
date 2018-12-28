package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/octowhale/iPicka/backend"
	"github.com/octowhale/iPicka/storage"
	"github.com/octowhale/iPicka/util"
	log "github.com/qiniu/x/log.v7"
	"github.com/sirupsen/logrus"
)

func Once(file string, HTTPSSchema string) {

	fileUrl, err := Upload(file)
	if err != nil {
		// logrus.Errorln(err)
		panic(err)
	}

	log.Infof("fileUrl: %s", fileUrl)
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

func SetDB(fileUrl, file string) error {

	md5sum, err := util.GetMd5(file)
	if err != nil {
		return err
	}

	backendClient, _ := backend.New(config.Backend)

	_, err = backendClient.Set(md5sum, fileUrl)
	if err != nil {
		return err
	}

	s, _ := backendClient.Get(md5sum)
	fmt.Println(s)

	return nil
}
