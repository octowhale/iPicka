package main

// https://github.com/mkideal/cli

import (
	"fmt"
	"log"
	"os"

	"github.com/mkideal/cli"
	"github.com/octowhale/iPicka/command"
)

func main() {

	if err := cli.Root(root,
		cli.Tree(put),
		cli.Tree(configure)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// root command
type rootT struct {
	cli.Helper
	// Name string `cli:"name" usage:"your name"`
}

var root = &cli.Command{
	Desc: "this is root command",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		// argv := ctx.Argv().(*rootT)
		// ctx.String("Hello, root command, I am %s\n", argv.Name)
		return nil
	},
}

// put command
type childPutArgv struct {
	cli.Helper
	File    string `cli:"file,f" usage:"指定上传文件"`
	Profile string `cli:"profile,p" usage:"指定图床配置文件, (default)" dft:"default" `
}

var put = &cli.Command{
	Name: "put",
	Desc: "上传文件",
	Argv: func() interface{} { return new(childPutArgv) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childPutArgv)
		profile := command.ProfileLoader(argv.Profile)
		command.OSSMain(profile, argv.File)
		return nil
	},
}

// configure command
type childConfigureArgv struct {
	cli.Helper
	Profile string `cli:"profile" usage:"指定图床配置文件, (default)" dft:"default" `
	// Brofile string `cli:"brofile" usage:"指定图床配置文件, (default)" dft:"default" `
	// Export  string `cli:"export" usage:"导出配置模板"`
}

var configure = &cli.Command{
	Name: "configure",
	Desc: "参数管理",
	Argv: func() interface{} { return new(childConfigureArgv) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childConfigureArgv)
		log.Println(argv)
		// log.Println(argv.Brofile)
		command.Configure()
		return nil
	},
}
