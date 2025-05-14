package mucl

import "github.com/iancoleman/strcase"

func (m *Config) FileName() string {
	if m == nil {
		return ""
	}
	return strcase.ToSnake(m.Name) + ".go"
}

func (m *Config) Fields() []*Field {
	if m == nil {
		return nil
	}
	var fields []*Field
	for _, entry := range m.Entries {
		if entry.Field != nil {
			fields = append(fields, entry.Field)
		}
	}
	return fields
}

func (m *Config) Configs() []*Config {
	if m == nil {
		return nil
	}
	var msgs []*Config
	for _, entry := range m.Entries {
		if entry.Config != nil {
			msgs = append(msgs, entry.Config)
		}
	}
	return msgs
}

func (m *Config) Enums() []*Enum {
	if m == nil {
		return nil
	}
	var enums []*Enum
	for _, entry := range m.Entries {
		if entry.Enum != nil {
			enums = append(enums, entry.Enum)
		}
	}
	return enums
}
