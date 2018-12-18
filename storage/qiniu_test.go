package storage

import (
	"os"
	"testing"
)

func Test_Qiniu(t *testing.T) {

	qiniu := Qiniu{}

	qiniu.key = os.Getenv("QINIU_ROOT_ACC_KEY")
	qiniu.sec = os.Getenv("QINIU_ROOT_ACC_SEC")
	qiniu.bucket = "filecdn"
	qiniu.region = "huadong"
	qiniu.customDomain = "cdn.tangx.in"

	Logger().Infoln(qiniu)

	qiniu.Put("naruto1.jpg", "/data/tmp/naruto.jpg")
}
