package command

import "github.com/mkideal/cli"

// root command
type rootT struct {
	cli.Helper
	// Conf string `cli:"conf,c" usage:"ipicka config file" dft:"$HOME/.ipika.json"`
}

var root = &cli.Command{
	Desc: "iPicka is a image docker agent by golang",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		return nil
	},
}
