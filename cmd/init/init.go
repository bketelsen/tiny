// Package init provides the init command for creating a new service
package init

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bketelsen/tiny/cmd"
	"github.com/bketelsen/tiny/project"

	"github.com/charmbracelet/huh"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:   "init",
			Usage:  "create an initial tiny configuration",
			Action: Run,
			Flags:  Flags,
		},
	)
}

func Run(c *cli.Context) error {
	var serviceName string
	var endpointName string
	var rpcMethodName string
	err := huh.NewForm(
		huh.NewGroup(

			huh.NewInput().Title("Service Name").
				Placeholder("Name of the service. e.g. users").
				Value(&serviceName).
				Validate(func(s string) error {
					if s == "" {
						return errors.New("service name cannot be empty")
					}
					if strings.Contains(s, " ") {
						return errors.New("service name cannot contain spaces")
					}
					return nil
				}),
			huh.NewInput().Title("Endpoint name").
				Placeholder("Name of the Endpoint. e.g. User").
				Value(&endpointName).
				Validate(func(s string) error {
					if s == "" {
						return errors.New("endpoint name cannot be empty")
					}
					if strings.Contains(s, " ") {
						return errors.New("endpoint name cannot contain spaces")
					}
					return nil
				}),
			huh.NewInput().Title("Method name").
				Placeholder("Name of the first method. e.g. Create").
				Value(&rpcMethodName).
				Validate(func(s string) error {
					if s == "" {
						return errors.New("method name cannot be empty")
					}
					if strings.Contains(s, " ") {
						return errors.New("method name cannot contain spaces")
					}
					return nil
				}),
		).Title("Create a new tiny service definition"),
	).
		WithTheme(huh.ThemeBase()).Run()
	if err != nil {
		if err == huh.ErrUserAborted {
			os.Exit(130)
		}
		fmt.Println(err)
		os.Exit(1)
	}

	// check to see if the current directory is empty
	// if not, warn the user
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	if len(files) > 0 {
		fmt.Println("Current directory is not empty")
		fmt.Println("Creating a new directory: ", serviceName)
		err = os.Mkdir(serviceName, 0o755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = os.Chdir(serviceName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Create the tiny file
	err = project.CreateConfig(serviceName, endpointName, rpcMethodName, serviceName, c.String("definition"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Created %s\n", c.String("definition"))
	fmt.Println("Edit the file and run `tiny gen` to generate the code")
	return nil
}

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "definition",
		Usage:   "tiny definition file",
		EnvVars: []string{"TINY_DEFINITION"},
		Value:   "service.tiny",
	},
}
