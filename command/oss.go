package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func OSSPutObject(bucket *oss.Bucket, srcPath string, destPath string) (err error) {
	err = bucket.PutObjectFromFile(destPath, srcPath)

	return
}

func OSSGetObjectURL(bucket *oss.Bucket) {

}

func OSSBucketClient(accKeyID string, accKeySecret string, bucketName string, endpoint string) (bucket *oss.Bucket, err error) {
	client, err := oss.New(endpoint, accKeyID, accKeySecret)
	if err != nil {
		handleError(err)
	}

	// 使用 bucket
	bucket, err = client.Bucket(bucketName)

	return
}

type OSSConfigStruct struct {
	AccKeyID     string `json:"accKeyID"`
	AccKeySecret string `json:"accKeySecret"`
	BucketName   string `json:"bucketName"`
	Domain       string `json:"domain,omitempty"`
	Endpoint     string `json:"endpoint"`
}

func OSSmain(srcPath string) {

	configFile := "config/ipicka.json"

	// open file
	ctx, _ := ioutil.ReadFile(configFile)

	// json unmarshal
	var config OSSConfigStruct
	err := json.Unmarshal(ctx, &config)

	if err != nil {

	}

	// check file is or not exist
	// srcPath := os.Args[1]
	_, err = os.Stat(srcPath)
	if os.IsNotExist(err) {
		log.Fatal("ERROR: target is not exist!")
	}

	Date := utils.DateF()
	destPath := Date + "/" + path.Base(srcPath)
	// get bucket client
	bucket, _ := OSSBucketClient(config.AccKeyID, config.AccKeySecret, config.BucketName, config.Endpoint)

	// upload object
	err = OSSPutObject(bucket, srcPath, destPath)

	if err != nil {
		log.Println(err)
	}

	var url string
	if config.Domain == "" {
		url = "http://" + config.BucketName + "." + config.Endpoint + "/" + destPath
	} else {
		url = config.Domain + "/" + destPath
	}
	// log.Println(url)
	utils.Output(url)
}
