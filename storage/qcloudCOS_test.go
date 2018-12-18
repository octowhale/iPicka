package storage

import (
	"os"
	"testing"
)

func Test_COS(t *testing.T) {

	cos := &QcloudCOS{}

	cos.key = os.Getenv("QCLOUD_ROOT_ACC_KEY")
	cos.sec = os.Getenv("QCLOUD_ROOT_ACC_SEC")
	cos.endpoint = "https://tangxin-test-02-1251448501.cos.ap-chengdu.myqcloud.com"

	var ipic Iipicka
	ipic = cos
	ipic.Put("naruto2.jpg", "/data/tmp/naruto.jpg")
}
