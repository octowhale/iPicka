package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/sirupsen/logrus"

	log "github.com/octowhale/iPicka/log"
	"github.com/octowhale/iPicka/util"
)

// https://github.com/mkideal/cli

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	flag.Parse()
	// if err := initConfig(); err != nil {
	// 	log.Fatal(err.Error())
	// }

	// client := storage.New(config)
	// client.Put("test/naruto", "/data/tmp/naruto.jpg")
	// client, err := storage.New(config.StorageConfig)

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	filepath := "/data/tmp/naruto.jpg"

	// check file already uploaded
	// s, err := CheckDatabase(filepath)
	fileMD5 := util.GetMd5(filepath)

	s, err := backend.Get(fileMD5)
	if err == nil {
		logrus.Debug("file already exist")
		fmt.Println(s)
	} else {
		logrus.Debug("file is not exist")
		var customPath string
		if len(config.Prefix) == 0 {
			customPath = fmt.Sprintf("https://%s/%s", config.CustomDomain, path.Base(filepath))
		} else {
			customPath = fmt.Sprintf("https://%s/%s/%s", config.CustomDomain, config.Prefix, path.Base(filepath))
		}
		backend.Set(fileMD5, customPath)
		// backend .Get(fileMD5)
		object := util.FullObject(config.Prefix, filepath)

		err = client.Put(object, filepath)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}
