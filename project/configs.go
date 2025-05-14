package project

import "github.com/iancoleman/strcase"

type Config struct {
	Name     string
	FieldMap map[string]*Field
	Options  Options
}

func NewConfig(name string) *Config {
	return &Config{
		Name:     name,
		FieldMap: make(map[string]*Field),
		Options:  make(Options),
	}
}

func (m *Config) GetField(name string) (*Field, bool) {
	field, ok := m.FieldMap[name]
	return field, ok
}

func (m *Config) GetAllFields() []*Field {
	fields := make([]*Field, 0, len(m.FieldMap))
	for _, field := range m.FieldMap {
		fields = append(fields, field)
	}
	return fields
}

func (m *Config) GetFieldNames() []string {
	fieldNames := make([]string, 0, len(m.FieldMap))
	for name := range m.FieldMap {
		fieldNames = append(fieldNames, name)
	}
	return fieldNames
}

func (m *Config) GetFieldTypes() []string {
	fieldTypes := make([]string, 0, len(m.FieldMap))
	for _, field := range m.FieldMap {
		fieldTypes = append(fieldTypes, field.TypeName)
	}
	return fieldTypes
}

func (m *Config) StructTag() string {
	if m == nil {
		return ""
	}
	return "`json:\"" + strcase.ToSnake(m.Name) + "\" yaml:\"" + strcase.ToSnake(m.Name) + "\"`"
}
