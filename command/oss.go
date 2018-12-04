package command

import (
	"fmt"

	"github.com/mkideal/cli"
)

type ossT struct {
	cli.Helper
}

var oss = &cli.Command{
	Name: "oss",
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
	fmt.Println("Hello, ", argv)
}
