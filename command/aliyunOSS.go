package command

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// AliyunOSS is the config struct of Qcloud OSS bucket
type AliyunOSS struct {
	key          string
	secret       string
	endpoint     string
	bucket       string
	customDomain string
}

// upload put file to remote bucket
func (ali *AliyunOSS) upload(bucket *oss.Bucket, objectKey string, filepath string) (bool, error) {
	// bucket := ali.load()

	err := bucket.PutObjectFromFile(objectKey, filepath)
	if err != nil {
		Logger().Errorln(err)
		return false, err
	}
	return true, err
}

// load return Aliyun Bucket Connection
func (ali *AliyunOSS) load() (bucket *oss.Bucket) {

	client, err := oss.New(ali.endpoint, ali.key, ali.secret)
	if err != nil {
		Logger().Error(err)
	}
	Logger().Debug(client)

	bucket, err = client.Bucket(ali.bucket)
	if err != nil {
		Logger().Error(err)
	}
	Logger().Debug(bucket)

	return bucket
}

// Put file to remote
func (ali *AliyunOSS) Put(objectKey string, filepath string) {
	bucket := ali.load()
	ok, _ := ali.upload(bucket, objectKey, filepath)

	var domain string
	if ali.customDomain != "" {
		domain = "https://" + ali.customDomain
	} else {
		domain = "https://" + ali.bucket + "." + ali.endpoint
	}

	if ok {
		fmt.Println("![](" + domain + "/" + objectKey + ")")
	}
}
