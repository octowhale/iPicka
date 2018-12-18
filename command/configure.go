package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mkideal/cli"
)

const (
	configPath = "config/config.json"
)

func configReader(configPath string) (j ConfigureT) {
	b, _ := ioutil.ReadFile(configPath)
	// j = ConfigureT{}
	json.Unmarshal(b, &j)

	return j
}

type ConfigureT struct {
	cli.Helper
	Provider     string `cli:"provider" usage:"Provider" prompt:"Provider[aliyunoss/qcloudcos/qiniu] "`
	Key          string `cli:"key" usage:"Bucket Access Key ID" prompt:"Access Key ID "`
	Sec          string `cli:"sec" usage:"Bucket Access Key Secret" prompt:"Access Key Secret "`
	Bucket       string `cli:"bucket" usage:"Bucket Name" prompt:"Bucket Name "`
	Endpoint     string `cli:"endpoint" usage:"Bucket Endpoint" prompt:"Bucket Endpoint "`
	Region       string `cli:"region" usage:"Region" prompt:"Bucket Region "`
	Schema       string `cli:"schema" usage:"Schema" prompt:"Schema[http/https] "`
	CustomDomain string `cli:"customdomain" usage:"Custom Domain" prompt:"Custom Domain "`
	Prefix       string `cli:"prefix" usage:"objectKey prefix" prompt:"objectKey Prefix "`
}

var configure = &cli.Command{
	Name: "configure",
	Desc: "configure your profile",
	Argv: func() interface{} { return new(ConfigureT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*ConfigureT)
		// ctx.String("Hello, child command, I am %s\n", argv.Profile)
		dumpConfig(argv)
		return nil
	},
}

func dumpConfig(argv *ConfigureT) {
	fmt.Println(argv)

	// c := new(Config)
}
