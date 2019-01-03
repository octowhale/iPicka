package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/octowhale/iPicka/util"
	"github.com/sirupsen/logrus"
)

func Once(file string, HTTPSSchema string) {

	logrus.Debugf("Entering Once")

	// md5sum, _ := util.GetMd5(file)
	// check exists
	// s, _ := backend.Get(md5sum)
	fileURL, err := Upload(file)
	if err != nil {
		logrus.Errorln(err)
		// panic(err)
	}

	logrus.Debugf("fileURL: %s", HTTPSSchema+fileURL)
	s, _ := SetDB(HTTPSSchema+fileURL, file)

	fmt.Sprintf("![](%s)", s)

}

func Upload(file string) (string, error) {

	logrus.Debugf("Entering Upload")
	// storageClient, _ := storage.New(config.Storage)

	if ok, err := util.IsFileExist(file); !ok {
		return "", err
	}
	logrus.Debugf("Files Exist check done!")

	if ok, err := util.IsSymlink(file); !ok {
		return "", err
	}
	logrus.Debugf("Files symblink check done!")

	var object string
	if len(config.Storage.Prefix) != 0 {
		object = strings.Trim(config.Storage.Prefix, "/") + "/" + path.Base(file)
	} else {
		object = path.Base(file)
	}
	logrus.Debugf("object combination done: object = %v", object)

	fmt.Printf("dafjlaksdjfalsdfj %v\n", config.Backend)
	fmt.Println("")
	var fileURL string
	// storageClient 设置为全局变量之后， 在这里不能使用。
	storageClient.Ping()
	fileURL, err := storageClient.Put(object, file)
	if len(config.Storage.CustomDomain) == 0 {
		// start to upload
		if err != nil {
			return "", err
		}
		return fileURL, nil
	}

	fileURL = config.Storage.CustomDomain + "/" + object
	logrus.Debugf("UPload: %s", fileURL)

	return fileURL, nil

}

func SetDB(fileURL, file string) (string, error) {

	md5sum, err := util.GetMd5(file)
	if err != nil {
		return "", err
	}

	// backendClient, _ := backend.New(config.Backend)

	url, err := backendClient.Get(md5sum)
	if err != nil {
		return "", err
	}
	if len(url) != 0 {
		// return url
		logrus.Debugln("Backend Get: ", url)
		return url, nil
	}

	_, err = backendClient.Set(md5sum, fileURL)
	if err != nil {
		return "", err
	}

	// url, _ = backendClient.Get(md5sum)
	// logrus.Debugln(url)

	// return url, nil

	return fileURL, nil
}

func DirMode(target, HTTPSSchema string) {
	files, _, _ := util.WalkDirectory(target)

	// logrus.Debugf("%v", files)
	for _, file := range files {
		Once(file, HTTPSSchema)
	}
}
