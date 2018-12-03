package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mkideal/cli"
)

type argConf struct {
	cli.Helper
	AccKeyID     string `cli:"accKeyID" usage:"ACCESS KEY for bucket" prompt:"ACCESS KEY"`
	AccKeySecret string `cli:"accKeySecret" usage:"ACCESS SECRET for bucket" prompt:"ACCESS SECRET"`
	BucketName   string `cli:"bucketName" usage:"bucket name" prompt:"BUCKET NAME"`
	Endpoint     string `cli:"endpoint" usage:"endpoint for your bucket" prompt:"BUCKET ENDPOINT"`
	Domain       string `cli:"domain" usage:"custom domain. Cname for bucket domain" prompt:"CUSTOM DOMAIN"`
}

func Configure() {
	configure()
}

func configure() {
	cli.Run(new(argConf), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argConf)
		ctx.String("username=%s, password=%s\n", argv.AccKeyID, argv.AccKeySecret)

		profile := map[string]string{"AccKeyID": argv.AccKeyID, "AccKeySecret": argv.AccKeySecret, "BucketName": argv.BucketName, "Endpoint": argv.Endpoint, "Domain": argv.Domain}

		data, _ := json.Marshal(profile)

		fmt.Printf("%s\n", data)

		ioutil.WriteFile("config/ipicka.json", data, 0644)
		return nil
	})
}

func json2file() {

}
