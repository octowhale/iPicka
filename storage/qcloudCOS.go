package storage

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/octowhale/iPicka/utils"
	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// QcloudCOS is the config struct of Qcloud OSS bucket
type QcloudCOS struct {
	Key          string
	Sec          string
	Endpoint     string
	Schema       string
	CustomDomain string
}

// Put do upload thing to cos bucket
func (qcloud *QcloudCOS) Put(objectKey string, filepath string) {
	qcloud.upload(objectKey, filepath)
}

func (qcloud *QcloudCOS) upload(objectKey string, filepath string) {

	u, err := url.Parse(qcloud.Endpoint)
	if err != nil {
		utils.Logger().Errorln(err)
	}

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

	fBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		utils.Logger().Errorln(err)
	}
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
	_, err = c.Object.Put(context.Background(), objectKey, f, nil)
	if err != nil {
		utils.Logger().Errorln(err)
		panic(err)
	}
}

func (qcloud *QcloudCOS) load() {

}
