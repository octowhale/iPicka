package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mkideal/cli"
	"github.com/octowhale/iPicka/utils"
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

// ConfigureT is the config os project
type ConfigureT struct {
	cli.Helper
	Profile      string `cli:"profile" usage:"profile" dft:"default" prompt:"profile "`
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
	// s, err := json.Marshal(argv)
	s, err := json.MarshalIndent(argv, "", "  ")
	fmt.Println(string(s))
	if err != nil {
		utils.Logger().Errorln(err)
		os.Exit(1)
	}
	flag := os.O_CREATE | os.O_TRUNC | os.O_RDWR
	fobj, err := os.OpenFile(configPath, flag, 0644)
	if err != nil {
		utils.Logger().Errorln(err)
		os.Exit(1)
	}
	defer fobj.Close()

	fobj.WriteString(string(s))
}
