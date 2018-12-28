package aliyunoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AccKey     string
	AccSec     string
	BucketName string
	Region     string
	Internal   bool
	Endpoint   string
	bucket     *oss.Bucket
}

func NewOSSClient(key, sec, bucketname, region string, internal bool) *Config {

	return &Config{
		AccKey:     key,
		AccSec:     sec,
		Region:     region,
		BucketName: bucketname,
		Internal:   internal,
	}
}

func (ali *Config) getClient() (client *oss.Client, err error) {

	var endpoint string
	// oss-cn-hangzhou-internal.aliyuncs.com
	if ali.Internal {
		endpoint = "https://oss-cn-hangzhou-internal.aliyuncs.com"
	} else {
		endpoint = "https://oss-cn-hangzhou.aliyuncs.com"
	}

	client, err = oss.New(endpoint, ali.AccKey, ali.AccSec)
	if err != nil {
		logrus.Errorln("oss.New() Error:", err)
		panic(err)
		// return nil, err
	}

	return client, err

}

func (ali *Config) getBucket() (*oss.Bucket, error) {

	if ali.bucket != nil {
		logrus.Debugln(ali.bucket.BucketName)
	} else {

		client, err := ali.getClient()
		if err != nil {
			logrus.Errorln(err)
		}

		ali.bucket, err = client.Bucket(ali.BucketName)
		if err != nil {
			logrus.Debugln("oss.Client():", err)
			return nil, err
		}
	}

	return ali.bucket, nil

}

func (ali *Config) Put(object, file string) (err error) {

	bucket, err := ali.getBucket()
	if err != nil {
		logrus.Debugln("Get bucket Error:", err)
		panic(err)
	}

	err = bucket.PutObjectFromFile(object, file)
	if err != nil {
		logrus.Debugln("Put File:", err)
		return err
	}

	return nil
}
