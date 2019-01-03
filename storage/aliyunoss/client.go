package aliyunoss

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AccKey       string
	AccSec       string
	BucketName   string
	Region       string
	Internal     bool
	Endpoint     string
	bucket       *oss.Bucket
	CustomDomain string
	Prefix       string
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

	// var endpoint string
	// oss-cn-hangzhou-internal.aliyuncs.com
	if ali.Internal {
		// ali.Endpoint = "https://oss-cn-hangzhou-internal.aliyuncs.com"
		ali.Endpoint = "oss-cn-hangzhou-internal.aliyuncs.com"
	} else {
		// ali.Endpoint = "https://oss-cn-hangzhou.aliyuncs.com"
		ali.Endpoint = "oss-cn-hangzhou.aliyuncs.com"
	}

	client, err = oss.New(ali.Endpoint, ali.AccKey, ali.AccSec)
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

func (ali *Config) Put(object, file string) (fileurl string, err error) {

	logrus.Debugf("Entering Aliyun oss Put")
	bucket, err := ali.getBucket()
	if err != nil {
		logrus.Errorln("Get bucket Error:", err)
		panic(err)
	}
	logrus.Debugf("Aliyun oss Bucket Loaded sucess.")

	err = bucket.PutObjectFromFile(object, file)
	if err != nil {
		logrus.Errorln("Put File: %v", err)
		return "", err
	}

	logrus.Debugf("aliyun oss put: %s.%s/%s", ali.BucketName, ali.Endpoint, object)
	return fmt.Sprintf("%s.%s/%s", ali.BucketName, ali.Endpoint, object), nil
}

// func (ali *Config) Upload(file string) (string, error) {

// 	// ali.Endpoint
// 	// if len(ali.CustomDomain) == 0 {
// 	// 	ali.CustomDomain = ali.Endpoint
// 	// }

// 	object := ali.Prefix + "/" + path.Base(file)

// 	_, err := ali.Put(object, file)
// 	if err != nil {
// 		return "", err
// 	}
// 	s := fmt.Sprintf("%s/%s", object, object)
// 	return s, nil

// }

func (ali *Config) Ping() error {

	logrus.Debugf("Entering Aliyun oss Put")
	bucket, err := ali.getBucket()
	if err != nil {
		return err
	}
	// fmt.Println("hello")

	logrus.Debugf("aliyun Storage Ping %+v", bucket)
	// listBucketResult, err := bucket.ListObjects()
	// if err != nil {
	// 	panic(err)
	// }
	// logrus.Debugf("listBucketResult: %+v", listBucketResult.)
	return nil
}
