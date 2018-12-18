package command

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

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
	if err := cli.Root(root,
		cli.Tree(configure)).Run(os.Args[1:]); err != nil {
		fmt.Println(os.Stdout, err)
		os.Exit(1)
	}
}
