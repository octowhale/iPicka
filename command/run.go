package command

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	configPath = "config/config.json"
)

// ipickaer is the interface
type ipickaer interface {
	Put(objectKey string, filepath string)
}

// Do is the main entrance
func Do() {
	var ipic ipickaer
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		Logger().Errorln(err)
	}

	c := Config{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		Logger().Errorln(err)
	}

	switch c.Provider {
	case "qcloudcos":
		ipic = &QcloudCOS{c.Key, c.Sec, c.Endpoint, c.Schema, c.CustomDomain}
	case "qiniu":
		ipic = &Qiniu{c.Key, c.Sec, c.Bucket, c.Region, c.CustomDomain}
	default:
		ipic = &AliyunOSS{c.Key, c.Sec, c.Endpoint, c.Bucket, c.CustomDomain}
	}

	objectKey, filepath := os.Args[1], os.Args[2]
	ipic.Put(objectKey, filepath)
}
