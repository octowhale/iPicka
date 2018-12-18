package command

import (
	"context"
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// Qiniu is the config struct of Qcloud OSS bucket
type Qiniu struct {
	key          string // 密钥 key
	sec          string // 密钥 secret
	bucket       string // bucket 名称
	region       string // bucket 区域
	customDomain string // 用户自定域名
}

// Put 上传文件到 qiniu bucket
//
// objectKey 文件在 bucket 中的路径, 不以 / 开头, 例如 image/test.png
// filepath  文件在本地的路径。
func (qiniu *Qiniu) Put(objectKey string, filepath string) {
	qiniu.upload(objectKey, filepath)
}

func cfgZone(region string) *storage.Zone {
	zones := map[string]*storage.Zone{"": &storage.ZoneHuadong}
	return zones[region]
}

func (qiniu *Qiniu) upload(key string, localFile string) {

	Logger().Debugln(qiniu)
	bucket := qiniu.bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(qiniu.key, qiniu.sec)
	Logger().Debugln(mac)

	upToken := putPolicy.UploadToken(mac)
	Logger().Debugln(upToken)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// cfg.Zone = cfgZone(qiniu.region)
	Logger().Debugln(cfg.Zone)

	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		Logger().Errorln(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)

	// fmt.Println(len(qiniu.customDomain))
	if len(qiniu.customDomain) != 0 {
		fmt.Println("![](https://" + qiniu.customDomain + "/" + ret.Key + ")")
	}

}
