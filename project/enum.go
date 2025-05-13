package project

type Enum struct {
	Name   string
	Values []*KeyValue
}

type KeyValue struct {
	Key   string
	Value int
}

func (e *Enum) GetKeys() []string {
	keys := make([]string, 0, len(e.Values))
	for _, kv := range e.Values {
		keys = append(keys, kv.Key)
	}
	return keys
}
