package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/bketelsen/tiny/project/templates"
)

func (p *Project) GenerateGoMod() error {
	cmd := exec.Command("go", "mod", "init", p.Module)
	cmd.Dir = p.outDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run go mod init: %v", err)
	}

	return nil
}

func (p *Project) Tidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = p.outDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run go mod tidy: %v", err)
	}

	return nil
}

func (p *Project) Format() error {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = p.outDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run go fmt: %v", err)
	}

	return nil
}

func (p *Project) GenerateTypes() error {
	typePath := p.outDir
	// Check if the directory exists, if not create it
	if _, err := os.Stat(typePath); os.IsNotExist(err) {
		err := os.MkdirAll(typePath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	typeFile, err := os.Create(fmt.Sprintf("%s/%s", typePath, "types.go"))
	if err != nil {
		return err
	}
	defer typeFile.Close()

	typeTemplate := template.Must(template.New("type").Parse(string(templates.TypeTemplate())))
	err = typeTemplate.Execute(typeFile, p)
	if err != nil {
		return err
	}

	configFile, err := os.Create(fmt.Sprintf("%s/%s", typePath, "config.go"))
	if err != nil {
		return err
	}
	defer configFile.Close()

	configTemplate := template.Must(template.New("configs").Parse(string(templates.ConfigurationTemplate())))
	err = configTemplate.Execute(configFile, p)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GenerateServers() error {
	serverPath := filepath.Join(p.outDir, "cmd", p.Service.DirectoryName())
	// Check if the directory exists, if not create it
	if _, err := os.Stat(serverPath); os.IsNotExist(err) {
		err := os.MkdirAll(serverPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	serverFilePath := fmt.Sprintf("%s/%s", serverPath, "main.go")
	if p.safeWriteFile(serverFilePath) {
		serverFile, err := os.Create(serverFilePath)
		if err != nil {
			return err
		}
		defer serverFile.Close()

		serverTemplate := template.Must(template.New("type").Parse(string(templates.ServiceTemplate())))
		err = serverTemplate.Execute(serverFile, map[string]interface{}{
			"Service": p.Service,
			"Module":  p.Module,
		})
		if err != nil {
			return err
		}
	}

	clientPath := p.outDir
	clientFilePath := fmt.Sprintf("%s/%s", clientPath, "client.go")
	if p.safeWriteFile(clientFilePath) {
		clientFile, err := os.Create(clientFilePath)
		if err != nil {
			return err
		}
		defer clientFile.Close()

		clientTemplate := template.Must(template.New("client").Parse(string(templates.ServiceClientTemplate())))
		err = clientTemplate.Execute(clientFile, map[string]interface{}{
			"Service": p.Service,
			"Module":  p.Module,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) GenerateHandlers() error {
	for _, t := range p.Service.GetAllEndpoints() {
		handlerPath := filepath.Join(p.outDir, "handlers")
		// Check if the directory exists, if not create it
		if _, err := os.Stat(handlerPath); os.IsNotExist(err) {
			err := os.MkdirAll(handlerPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		}
		handlerFilePath := fmt.Sprintf("%s/%s", handlerPath, strings.ToLower(t.Name)+".go")
		if p.safeWriteFile(handlerFilePath) {
			handlerFile, err := os.Create(fmt.Sprintf("%s/%s", handlerPath, strings.ToLower(t.Name)+".go"))
			if err != nil {
				return err
			}
			defer handlerFile.Close()

			handlerTemplate := template.Must(template.New("type").Parse(string(templates.HandlerTemplate())))
			err = handlerTemplate.Execute(handlerFile, map[string]interface{}{
				"Endpoint": t,
				"Module":   p.Module,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Project) GenerateTaskfile() error {
	taskfilePath := filepath.Join(p.outDir, "Taskfile.yml")

	if p.safeWriteFile(taskfilePath) {
		taskFile, err := os.Create(taskfilePath)
		if err != nil {
			return err
		}
		defer taskFile.Close()

		infraTemplate := template.Must(template.New("type").Parse(string(templates.TaskfileTemplate())))
		err = infraTemplate.Execute(taskFile,
			map[string]interface{}{
				"SERVICE_NAME": p.Service.DirectoryName(),
			})
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) GenerateGitIgnore() error {
	ignoreFilePath := filepath.Join(p.outDir, ".gitignore")

	if p.safeWriteFile(ignoreFilePath) {
		ignoreFile, err := os.Create(ignoreFilePath)
		if err != nil {
			return err
		}
		defer ignoreFile.Close()

		ignoreTemplate := template.Must(template.New("type").Parse(string(templates.GitIgnoreTemplate())))
		err = ignoreTemplate.Execute(ignoreFile,
			map[string]interface{}{
				"SERVICE_NAME": p.Service.DirectoryName(),
			})
		if err != nil {
			return err
		}
	}
	return nil
}
