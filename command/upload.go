package command

import (
	"path"

	"github.com/octowhale/iPicka/storage"
)

// ipickaer is the interface
type ipickaer interface {
	Put(objectKey string, filepath string)
}

// Do is the main entrance
func upload(filepath string) {
	var ipic ipickaer
	// b, err := ioutil.ReadFile(configPath)
	// if err != nil {
	// 	utils.Logger().Errorln(err)
	// }

	// // c := Config{}
	// c := ConfigureT{}
	// err = json.Unmarshal(b, &c)
	// if err != nil {
	// 	utils.Logger().Errorln(err)
	// }
	c := configReader(configPath)

	switch c.Provider {
	case "qcloudcos":
		ipic = &storage.QcloudCOS{c.Key, c.Sec, c.Endpoint, c.Schema, c.CustomDomain}
	case "qiniu":
		ipic = &storage.Qiniu{c.Key, c.Sec, c.Bucket, c.Region, c.CustomDomain}
	default:
		ipic = &storage.AliyunOSS{c.Key, c.Sec, c.Endpoint, c.Bucket, c.CustomDomain}
	}

	// objectKey, filepath := os.Args[1], os.Args[2]
	objectKey := c.Prefix + path.Base(filepath)
	ipic.Put(objectKey, filepath)

}
