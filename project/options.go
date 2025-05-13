package project

import "github.com/bketelsen/tiny/mucl"

// Options is a map of option names to their values.
type Options map[string]*Option

// Get returns the value of the option with the given name.
func (o Options) Get(name string) (*Option, bool) {
	option, ok := o[name]
	return option, ok
}

// GetAll returns all options.
func (o Options) GetAll() []*Option {
	options := make([]*Option, 0, len(o))
	for _, option := range o {
		options = append(options, option)
	}
	return options
}

// GetNames returns the names of all options.
func (o Options) GetNames() []string {
	names := make([]string, 0, len(o))
	for name := range o {
		names = append(names, name)
	}
	return names
}

type Option struct {
	Name  string
	Attr  *string
	Value Valuer
}

func NewOption(name string, attr *string, value *mucl.Value) *Option {
	return &Option{
		Name:  name,
		Attr:  attr,
		Value: parseValue(value),
	}
}

func parseValue(value *mucl.Value) Valuer {
	if value.String != nil {
		return &StringValue{Value: *value.String}
	} else if value.Number != nil {
		return &NumberValue{Value: *value.Number}
	} else if value.Int != nil {
		return &IntValue{Value: *value.Int}
	} else if value.Bool != nil {
		return &BoolValue{Value: *value.Bool}
	}

	return &StringValue{Value: "unknown value"} // Default case, should not happen
}
