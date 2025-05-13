// Package gen provides the gen command for generating code from a mucl file
package gen

import (
	"errors"
	"fmt"

	"github.com/bketelsen/tiny/cmd"
	"github.com/bketelsen/tiny/project"

	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:   "gen",
			Usage:  "generate code from a mucl file",
			Action: Run,
			Flags:  Flags,
		},
	)
}

func Run(c *cli.Context) error {
	file := c.String("definition")
	if file == "" {
		return errors.New("definition file is required")
	}
	onlyTypes := c.Bool("types")
	force := c.Bool("force")

	p, err := project.NewProject(
		project.WithMuclFile(file),
		project.WithOutputDir("."),
		project.WithOnlyTypes(onlyTypes),
		project.WithForce(force),
	)
	if err != nil {
		return fmt.Errorf("failed to create project: %v", err)
	}
	if err := p.Init(); err != nil {
		return fmt.Errorf("failed to initialize project: %v", err)
	}
	if err := p.Apply(); err != nil {
		return fmt.Errorf("failed to generate code: %v", err)
	}
	fmt.Println("Code generation complete")

	return nil
}

var Flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "types",
		Usage:   "only generate types",
		EnvVars: []string{"TINY_GEN_TYPES"},
	},
	&cli.BoolFlag{
		Name:    "force",
		Usage:   "destructively overwrite existing files",
		EnvVars: []string{"TINY_GEN_FORCE"},
	},
	&cli.StringFlag{
		Name:    "definition",
		Usage:   "tiny definition file",
		EnvVars: []string{"TINY_DEFINITION"},
		Value:   "service.tiny",
	},
}
