package project

type Valuer interface {
	// GetValue returns the value of the option.
	GetValue() interface{}
}

/*
type Value struct {
	Pos lexer.Position

	String    *string  `  @String`
	Number    *float64 `| @Float`
	Int       *int64   `| @Int`
	Bool      *bool    `| (@"true" | "false")`
	Reference *string  `| @Ident @( "." Ident )*`
	Map       *Map     `| @@`
	Array     *Array   `| @@`
}
*/

type StringValue struct {
	Value string
}

func (s *StringValue) GetValue() interface{} {
	return s.Value
}

type NumberValue struct {
	Value float64
}

func (n *NumberValue) GetValue() interface{} {
	return n.Value
}

type IntValue struct {
	Value int64
}

func (i *IntValue) GetValue() interface{} {
	return i.Value
}

type BoolValue struct {
	Value bool
}

func (b *BoolValue) GetValue() interface{} {
	return b.Value
}
