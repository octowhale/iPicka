package command

import (
	"os"
	"testing"
)

func Test_OSS(t *testing.T) {
	key := os.Getenv("ALIYUN_ROOT_ACC_KEY")
	sec := os.Getenv("ALIYUN_ROOT_ACC_SEC")
	endpoint := "oss-cn-hangzhou.aliyuncs.com"
	bucket := "tangxin-test-02"

	// oss := &AliyunOSS{key, sec, endpoint, bucket, ""}
	oss := &AliyunOSS{}
	oss.key = key
	oss.secret = sec
	oss.endpoint = endpoint
	oss.bucket = bucket

	var ipic Iipicka
	ipic = oss

	objectKey := "naruto5.jpg"
	filepath := "/data/tmp/naruto.jpg"
	ipic.Put(objectKey, filepath)
}
