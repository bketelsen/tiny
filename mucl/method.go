package mucl

import "github.com/iancoleman/strcase"

func (m *Method) FileName() string {
	if m == nil {
		return ""
	}
	return strcase.ToSnake(m.Name) + ".go"
}
