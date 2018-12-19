package command

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/octowhale/iPicka/storage"
	"github.com/octowhale/iPicka/utils"
)

// ipickaer is the interface
type ipickaer interface {
	Put(objectKey string, filepath string)
}

// Do is the main entrance
func upload(filepath string) {
	var ipic ipickaer
	c := configReader(configPath)

	switch c.Provider {
	case "qcloudcos":
		ipic = &storage.QcloudCOS{c.Key, c.Sec, c.Endpoint, c.Schema, c.CustomDomain}
	case "qiniu":
		ipic = &storage.Qiniu{c.Key, c.Sec, c.Bucket, c.Region, c.CustomDomain}
	default:
		ipic = &storage.AliyunOSS{c.Key, c.Sec, c.Endpoint, c.Bucket, c.CustomDomain}
	}

	// path exists
	fi, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		utils.Logger().Errorln(err)
		os.Exit(1)
	}

	// path is dir
	if fi.IsDir() {
		files, _, _ := fileAndDir(filepath)
		for _, subFilepath := range files {
			objectKey := c.Prefix + path.Base(subFilepath)
			ipic.Put(objectKey, subFilepath)
		}
	} else {
		// objectKey, filepath := os.Args[1], os.Args[2]
		objectKey := c.Prefix + path.Base(filepath)
		ipic.Put(objectKey, filepath)
	}

}

func fileAndDir(dirname string) (files []string, dirs []string, err error) {

	dir, err := ioutil.ReadDir(dirname)
	// fmt.Println(dir)
	if err != nil {
		return nil, nil, err
	}

	pathSep := string(os.PathSeparator)

	for _, fi := range dir {
		// fmt.Println(pos)
		if fi.IsDir() {
			dirs = append(dirs, dirname+pathSep+fi.Name())
		} else {
			files = append(files, dirname+pathSep+fi.Name())
		}
	}
	// fmt.Println(files)
	return
}
