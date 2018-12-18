package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/mkideal/cli"
	"github.com/octowhale/iPicka/storage"
	"github.com/octowhale/iPicka/utils"
)

const (
	configPath = "config/config.json"
)

// ipickaer is the interface
type ipickaer interface {
	Put(objectKey string, filepath string)
}

// Do is the main entrance
func upload(filepath string) {
	var ipic ipickaer
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		utils.Logger().Errorln(err)
	}

	c := Config{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		utils.Logger().Errorln(err)
	}

	switch c.Provider {
	case "qcloudcos":
		ipic = &storage.QcloudCOS{c.Key, c.Sec, c.Endpoint, c.Schema, c.CustomDomain}
	case "qiniu":
		ipic = &storage.Qiniu{c.Key, c.Sec, c.Bucket, c.Region, c.CustomDomain}
	default:
		ipic = &storage.AliyunOSS{c.Key, c.Sec, c.Endpoint, c.Bucket, c.CustomDomain}
	}

	// objectKey, filepath := os.Args[1], os.Args[2]
	objectKey := c.Prefix + path.Base(filepath)
	ipic.Put(objectKey, filepath)

}

type rootT struct {
	cli.Helper
	File string `cli:"file,f" usage:"filepath to upload"`
}

var root = &cli.Command{
	Name: "root",
	Desc: "upload file",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		upload(argv.File)
		return nil
	},
}

// Do the project
func Do() {
	// fmt.Println(len(os.Args))
	if len(os.Args) < 3 {

		utils.Logger().Errorln("To few argumetns")
		os.Exit(1)
	}

	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Println(os.Stdout, err)
		os.Exit(1)
	}
}
