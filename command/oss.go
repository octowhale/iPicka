package command

import (
	"log"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/mkideal/cli"
	"github.com/octowhale/iPicka/utils"
)

type ossT struct {
	cli.Helper
	Profile string `cli:"profile,p" usage:"image dock profile" dft:"default"`
	File    string `cli:"file,f" usage:"file to upload"`
}

var put = &cli.Command{
	Name: "put",
	Desc: "Put image to Aliyun OSS",
	Argv: func() interface{} { return new(ossT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*ossT)
		// ctx.String("Hello, child command, I am %s\n", argv.Profile)
		ossCommand(argv)
		return nil
	},
}

func ossCommand(argv *ossT) {
	// fmt.Println("Hello, ", argv)

	profile := getProfile(argv.Profile)
	bucket := ossClient(profile)

	objectKey := utils.DateF() + "/" + path.Base(argv.File)
	ossUpload(bucket, argv.File, objectKey)

	// output
	var url string
	if profile.Domain == "" {
		url = "https://" + profile.BucketName + "." + profile.Endpoint + "/" + objectKey
	} else {
		url = profile.Domain + "/" + objectKey
	}

	utils.Output(url)

}

func getProfile(key string) (profile Profile) {
	profile = ProfileLoader(key)
	// fmt.Println(profile)
	return
}

func ossClient(profile Profile) (bucket *oss.Bucket) {

	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := profile.Endpoint
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyID := profile.AccKeyID
	accessKeySecret := profile.AccKeySecret

	// 获取 client pointer
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Fatalln(err)
	}

	// 获取 bucket pointer
	bucket, err = client.Bucket(profile.BucketName)
	if err != nil {
		log.Fatalf("%s", err)
	}
	return
}

func ossUpload(bucket *oss.Bucket, filePath string, objectKey string) (err error) {

	err = bucket.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		log.Fatalf("%s", err)
	}

	return
}
