package command

import (
	"fmt"

	"github.com/mkideal/cli"
)

type configureT struct {
	cli.Helper
	// Gender       string `cli:"gender"`
	Profile      string `cli:"profile" usage:"specify profile" dft:"default" prompt:"Profile Name      "`
	AccKeyID     string `cli:"accKeyID" usage:"Bucket Access Key ID"         prompt:"Access Key ID     "`
	AccKeySecret string `cli:"accKeySecret" usage:"Bucket Access Key Secret" prompt:"Access Key Secret "`
	BucketName   string `cli:"bucketName" usage:"Bucket Name" prompt:"Bucket Name "`
	Endpoint     string `cli:"endpoint" usage:"Bucket Endpoint" prompt:"Bucket Endpoint "`
	Domain       string `cli:"domain" usage:"Custom Domain" prompt:"Custom Domain "`
}

var configure = &cli.Command{
	Name: "configure",
	Desc: "configure your profile",
	Argv: func() interface{} { return new(configureT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*configureT)
		// ctx.String("Hello, child command, I am %s\n", argv.Profile)
		configureCommand(argv)
		return nil
	},
}

func configureCommand(argv *configureT) {
	fmt.Println("Hello, ", argv)
}
