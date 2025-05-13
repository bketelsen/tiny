package project

type Message struct {
	Name     string
	FieldMap map[string]*Field
	Options  Options
}

func NewMessage(name string) *Message {
	return &Message{
		Name:     name,
		FieldMap: make(map[string]*Field),
		Options:  make(Options),
	}
}

func (m *Message) GetField(name string) (*Field, bool) {
	field, ok := m.FieldMap[name]
	return field, ok
}

func (m *Message) GetAllFields() []*Field {
	fields := make([]*Field, 0, len(m.FieldMap))
	for _, field := range m.FieldMap {
		fields = append(fields, field)
	}
	return fields
}

func (m *Message) GetFieldNames() []string {
	fieldNames := make([]string, 0, len(m.FieldMap))
	for name := range m.FieldMap {
		fieldNames = append(fieldNames, name)
	}
	return fieldNames
}

func (m *Message) GetFieldTypes() []string {
	fieldTypes := make([]string, 0, len(m.FieldMap))
	for _, field := range m.FieldMap {
		fieldTypes = append(fieldTypes, field.TypeName)
	}
	return fieldTypes
}
