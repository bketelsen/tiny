package project

import (
	"os"
	"text/template"

	"github.com/bketelsen/tiny/project/templates"
)

func CreateConfig(service, endpoint, method, module, file string) error {
	muFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer muFile.Close()

	configTemplate := template.Must(template.New("config").Parse(string(templates.ConfigTemplate())))
	err = configTemplate.Execute(muFile, map[string]interface{}{
		"Service":  service,
		"Endpoint": endpoint,
		"Method":   method,
		"Module":   module,
	})
	if err != nil {
		return err
	}
	return nil
}
