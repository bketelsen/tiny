package mucl

import "github.com/iancoleman/strcase"

func (m *Message) FileName() string {
	if m == nil {
		return ""
	}
	return strcase.ToSnake(m.Name) + ".go"
}

func (m *Message) Fields() []*Field {
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

func (m *Message) Messages() []*Message {
	if m == nil {
		return nil
	}
	var msgs []*Message
	for _, entry := range m.Entries {
		if entry.Message != nil {
			msgs = append(msgs, entry.Message)
		}
	}
	return msgs
}

func (m *Message) Enums() []*Enum {
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
