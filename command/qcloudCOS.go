package command

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

type QcloudCOS struct {
	key          string
	sec          string
	endpoint     string
	schema       string
	customDomain string
}

func (qcloud *QcloudCOS) Put(objectKey string, filepath string) {
	qcloud.upload(objectKey, filepath)
}

func (qcloud *QcloudCOS) upload(objectKey string, filepath string) {

	u, _ := url.Parse(qcloud.endpoint)

	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("QCLOUD_ROOT_ACC_KEY"),
			SecretKey: os.Getenv("QCLOUD_ROOT_ACC_SEC"),
			// Transport: &debug.DebugRequestTransport{
			// 	RequestHeader:  true,
			// 	RequestBody:    true,
			// 	ResponseHeader: true,
			// 	ResponseBody:   true,
			// },
		},
	})

	fBytes, _ := ioutil.ReadFile(filepath)
	// f := strings.NewReader("xxx")
	f := strings.NewReader(string(fBytes))

	// opt := &cos.ObjectPutOptions{
	// 	ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
	// 		ContentType: "text/html",
	// 	},
	// 	ACLHeaderOptions: &cos.ACLHeaderOptions{
	// 		XCosACL: "public-read",
	// 		// XCosACL: "private",
	// 	},
	// }
	_, err := c.Object.Put(context.Background(), objectKey, f, nil)
	if err != nil {
		panic(err)
	}
}

func (qcloud *QcloudCOS) load() {

}
