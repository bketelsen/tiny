// Package cmd provides the command line interface for mu
package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

type Cmd interface {
	// The cli app within this cmd
	App() *cli.App
	// Run executes the command
	Run() error
}

type command struct {
	app *cli.App
}

var (
	DefaultCmd Cmd = New()

	// name of the binary
	name = "tiny"
	// description of the binary
	description = "A microservice development tool for nats"
	repository  = "https://github.com/bketelsen/tiny"
	docs        = "https://bketelsen.github.io/tiny/"
	// defaultFlags which are used on all commands
	defaultFlags = []cli.Flag{}
)

func action(c *cli.Context) error {
	if c.Args().Len() == 0 {
		return MissingCommand(c)
	}

	// srv == nil
	return UnexpectedCommand(c)
}

func New() *command {
	cmd := new(command)
	cmd.app = cli.NewApp()
	cmd.app.Name = name
	cmd.app.Version = buildVersionEnhanced()
	cmd.app.Usage = description
	cmd.app.Flags = defaultFlags
	cmd.app.Action = action

	return cmd
}

func (c *command) App() *cli.App {
	return c.app
}

func (c *command) Run() error {
	return c.app.Run(os.Args)
}

// Register CLI commands
func Register(cmds ...*cli.Command) {
	app := DefaultCmd.App()
	app.Commands = append(app.Commands, cmds...)

	// sort the commands so they're listed in order on the cli
	// todo: move this to micro/cli so it's only run when the
	// commands are printed during "help"
	sort.Slice(app.Commands, func(i, j int) bool {
		return app.Commands[i].Name < app.Commands[j].Name
	})
}

// Run the default command
func Run() {
	if err := DefaultCmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
