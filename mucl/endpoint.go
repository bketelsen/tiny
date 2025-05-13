package mucl

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func (m *Endpoint) FileName() string {
	return "main.go"
}

func (m *Endpoint) DirectoryName() string {
	if m == nil {
		return ""
	}
	return strings.ToLower(m.Name)
}

func (m *Endpoint) ClientFileName() string {
	if m == nil {
		return ""
	}
	return strcase.ToSnake(m.Name) + "_client" + ".go"
}

func (m *Endpoint) ClientStructName() string {
	if m == nil {
		return ""
	}
	return strcase.ToLowerCamel(m.Name)
}

func (m *Endpoint) Methods() []*Method {
	if m == nil {
		return nil
	}
	var methods []*Method
	for _, entry := range m.Entry {
		if entry.Method != nil {
			methods = append(methods, entry.Method)
		}
	}
	return methods
}
