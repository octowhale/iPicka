package command

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/octowhale/iPicka/utils"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// OSSPutObject uploads file to OSS bucket
func OSSPutObject(bucket *oss.Bucket, srcPath string, destPath string) (err error) {
	err = bucket.PutObjectFromFile(destPath, srcPath)

	return
}

// OSSBucketClient returns bucket client
func OSSBucketClient(accKeyID string, accKeySecret string, bucketName string, endpoint string) (bucket *oss.Bucket, err error) {
	client, err := oss.New(endpoint, accKeyID, accKeySecret)
	if err != nil {
		handleError(err)
	}

	// 使用 bucket
	bucket, err = client.Bucket(bucketName)

	return
}

// OSSMain the entry of oss bucket loader
func OSSMain(profile Profile, src string) {

	// get bucket client
	bucket, _ := OSSBucketClient(profile.AccKeyID, profile.AccKeySecret, profile.BucketName, profile.Endpoint)

	// upload
	objectKey := utils.DateF() + "/" + path.Base(src)
	err := OSSPutObject(bucket, src, objectKey)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// output
	var url string
	if profile.Domain == "" {
		url = "https://" + profile.BucketName + "." + profile.Endpoint + "/" + objectKey
	} else {
		url = profile.Domain + "/" + objectKey
	}

	utils.Output(url)
}
