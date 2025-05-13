package main

import (
	"github.com/bketelsen/tiny/cmd"

	// load packages so they can register commands
	_ "github.com/bketelsen/tiny/cmd/ebnf"
	_ "github.com/bketelsen/tiny/cmd/gen"
	_ "github.com/bketelsen/tiny/cmd/init"
)

func main() {
	cmd.Run()
}
