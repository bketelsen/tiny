// Package project defines a tiny project structure
package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bketelsen/tiny/mucl"
)

// Project represents a project with a mucl
// definition file in a dedicated directory.
type Project struct {
	muFile    string
	outDir    string
	onlyTypes bool
	force     bool
	//	gen       *generator.Generator
	def     *mucl.Definition
	Module  string
	Service *Service
}

// ProjectOption is a function that configures a Project.
// It is used to set various options for the project.
type ProjectOption func(*Project)

// WithMuclFile sets the mucl file for the project.
func WithMuclFile(muclFile string) ProjectOption {
	return func(p *Project) {
		p.muFile = muclFile
	}
}

// WithOutputDir sets the output directory for the project.
// This is where the generated code will be placed.
func WithOutputDir(outDir string) ProjectOption {
	return func(p *Project) {
		p.outDir = outDir
	}
}

// WithOnlyTypes sets whether to only generate types.
func WithOnlyTypes(onlyTypes bool) ProjectOption {
	return func(p *Project) {
		p.onlyTypes = onlyTypes
	}
}

// WithForce sets whether to forcefully overwrite existing files.
func WithForce(force bool) ProjectOption {
	return func(p *Project) {
		p.force = force
	}
}

// NewProject creates a new Project instance with the given options.
// It initializes the project with default values and applies the provided options.
// It returns a pointer to the Project instance and an error if any occurs.
// The default mucl file is "service.mu" and the default output directory is ".".
// The project is not loaded at this point, so you need to call Init() separately.
func NewProject(opts ...ProjectOption) (*Project, error) {
	p := &Project{
		muFile: "service.tiny",
		outDir: ".",
	}
	for _, opt := range opts {
		opt(p)
	}
	return p, nil
}

func (p *Project) Init() error {
	if p.muFile == "" {
		return nil
	}
	if _, err := os.Stat(p.muFile); os.IsNotExist(err) {
		return fmt.Errorf("definition file does not exist: %s", p.muFile)
	}
	bb, err := os.ReadFile(p.muFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	def, err := mucl.Parser.ParseBytes(p.muFile, bb)
	if err != nil {
		return fmt.Errorf("parsing failure: %v", err)
	}

	p.def = def

	p.Service, err = fromMuCL(p.def)
	if err != nil {
		return fmt.Errorf("failed to parse service: %v", err)
	}
	return p.setGoModule()
}

// Apply applies the generator to the project.
func (p *Project) Apply() error {
	err := p.GenerateTypes()
	if err != nil {
		return fmt.Errorf("failed to generate types: %v", err)
	}
	if !p.onlyTypes {
		err := p.GenerateServers()
		if err != nil {
			return fmt.Errorf("failed to generate servers: %v", err)
		}
		err = p.GenerateHandlers()
		if err != nil {
			return fmt.Errorf("failed to generate handlers: %v", err)
		}
		err = p.GenerateTaskfile()
		if err != nil {
			return fmt.Errorf("failed to generate taskfile: %v", err)
		}
		err = p.GenerateGitIgnore()
		if err != nil {
			return fmt.Errorf("failed to generate .gitignore: %v", err)
		}
	}
	err = p.Format()
	if err != nil {
		return fmt.Errorf("failed to format code: %v", err)
	}
	return p.Tidy()
}

// setGoModule sets the Go module name for the project.
// It reads the go.mod file in the output directory and extracts the module name.
// If the go.mod file does not exist or cannot be read, it sets the module name to the output directory.
func (p *Project) setGoModule() error {
	var needsModule bool
	// Try to read go.mod file
	goModPath := p.computedPath("go.mod")

	data, err := os.ReadFile(goModPath)
	if err != nil {
		if os.IsNotExist(err) {
			needsModule = true
		} else {
			fmt.Printf("failed to read go.mod file: %v\n", err)
			return err
		}
	} else {
		needsModule = false
	}
	if !needsModule {
		// Parse the first line to extract module name
		lines := strings.SplitSeq(string(data), "\n")
		for line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "module ") {
				p.Module = strings.TrimSpace(strings.TrimPrefix(line, "module"))
				fmt.Printf("found module name: %s\n", p.Module)
				break
			}
		}
	} else {
		// if module name is empty, there is no go.mod file
		// set the module name to the name of the output directory
		if p.Module == "" {
			p.Module = p.outDir
		}
		if p.Module == "." {
			// if the output directory is ".", set the module name to the current directory
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Printf("failed to get current directory: %v\n", err)
				return err
			}
			p.Module = filepath.Base(pwd)
		}
	}
	if needsModule {
		err := p.GenerateGoMod()
		if err != nil {
			fmt.Printf("failed to generate go.mod file: %v\n", err)
			return err
		}
		fmt.Printf("generated go.mod file with module name: %s\n", p.Module)
	}
	return nil
}

func (p *Project) safeWriteFile(path string) bool {
	if p.force {
		return true
	}
	// Check if the file already exists
	if _, err := os.Stat(path); err == nil {
		// File exists, prompt the user for confirmation
		fmt.Println("*WARNING* Existing file detected.")
		fmt.Println("Use the --force flag to overwrite without confirmation.")
		fmt.Printf("File %s already exists. Overwrite? (y/n): ", path)
		var response string
		_, _ = fmt.Scanln(&response)
		return strings.ToLower(response) == "y"
	}
	return true
}
