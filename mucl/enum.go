package mucl

import "github.com/iancoleman/strcase"

func (m *Enum) FileName() string {
	if m == nil {
		return ""
	}
	return strcase.ToSnake(m.Name) + ".go"
}

func (m *EnumValue) ExportedName() string {
	if m == nil {
		return ""
	}
	if m.Key == "" {
		return ""
	}
	return strcase.ToCamel(m.Key)
}
