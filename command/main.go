package command

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

// Run start ipicka
func Run() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(configure),
		cli.Tree(put),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
