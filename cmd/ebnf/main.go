// Package gen provides the gen command for generating code from a mucl file
package gen

import (
	"fmt"

	"github.com/bketelsen/tiny/cmd"
	"github.com/bketelsen/tiny/mucl"

	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:   "ebnf",
			Usage:  "output EBNF grammar for the mucl file",
			Action: Run,
			Flags:  Flags,
			Hidden: true,
		},
	)
}

func Run(_ *cli.Context) error {
	fmt.Println(mucl.Parser.String())
	return nil
}

var Flags = []cli.Flag{}
