package command

import "testing"

var src = "/Users/tangxin/Desktop/back.jpg"

func Test_OSSMain(t *testing.T) {

	profile := ConfigLoader("aliyun")
	OSSMain(profile, src)
}
