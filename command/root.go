package command

import "github.com/mkideal/cli"

// root command
type rootT struct {
	cli.Helper
	// Name string `cli:"name" usage:"your name"`
}

var root = &cli.Command{
	Desc: "iPicka is a image docker agent by golang",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		// argv := ctx.Argv().(*rootT)
		// ctx.String("Hello, root command, I am %s\n", argv.Name)
		return nil
	},
}
