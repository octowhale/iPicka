package storage

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/octowhale/iPicka/utils"
)

// AliyunOSS is the config struct of Qcloud OSS bucket
type AliyunOSS struct {
	Key          string
	Secret       string
	Endpoint     string
	Bucket       string
	CustomDomain string
}

// upload put file to remote bucket
func (ali *AliyunOSS) upload(bucket *oss.Bucket, objectKey string, filepath string) (bool, error) {
	// bucket := ali.load()

	err := bucket.PutObjectFromFile(objectKey, filepath)
	if err != nil {
		utils.Logger().Errorln(err)
		return false, err
	}
	return true, err
}

// load return Aliyun Bucket Connection
func (ali *AliyunOSS) load() (bucket *oss.Bucket) {

	client, err := oss.New(ali.Endpoint, ali.Key, ali.Secret)
	if err != nil {
		utils.Logger().Error(err)
	}
	utils.Logger().Debug(client)

	bucket, err = client.Bucket(ali.Bucket)
	if err != nil {
		utils.Logger().Error(err)
	}
	utils.Logger().Debug(bucket)

	return bucket
}

// Put file to remote
func (ali *AliyunOSS) Put(objectKey string, filepath string) {
	bucket := ali.load()
	ok, _ := ali.upload(bucket, objectKey, filepath)

	var domain string
	if ali.CustomDomain != "" {
		domain = "https://" + ali.CustomDomain
	} else {
		domain = "https://" + ali.Bucket + "." + ali.Endpoint
	}

	if ok {
		fmt.Println("![](" + domain + "/" + objectKey + ")")
	}
}
