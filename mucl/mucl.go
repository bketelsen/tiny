// Package mucl provides a parser for the MuCL (Micro Configuration Language) protocol.
//
//nolint:govet
package mucl

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Definition struct {
	Pos lexer.Position

	Entries []*Entry `( @@ ";"* )*`
}

type Entry struct {
	Pos         lexer.Position
	Service     string    `  "service" "=" @String`
	Description string    `| "description" "=" @String`
	Message     *Message  `| @@`
	Endpoint    *Endpoint `| @@`
	Enum        *Enum     `| @@`
	Option      *Option   `| "option" @@`
	Config      *Config   `| @@`
}

type Config struct {
	Pos lexer.Position

	Name    string         `"config" @Ident`
	Entries []*ConfigEntry `"{" @@* "}"`
}
type ConfigEntry struct {
	Pos lexer.Position

	Enum   *Enum   `( @@`
	Config *Config ` | @@`
	Field  *Field  ` | @@ ) ";"*`
}

type Option struct {
	Pos lexer.Position

	Name  string  `( "(" @Ident @( "." Ident )* ")" | @Ident @( "." @Ident )* )`
	Attr  *string `( "." @Ident ( "." @Ident )* )?`
	Value *Value  `"=" @@`
}

type Value struct {
	Pos lexer.Position

	String *string  `  @String`
	Number *float64 `| @Float`
	Int    *int64   `| @Int`
	Bool   *bool    `| (@"true" | "false")`
}

type Service struct {
	Pos lexer.Position

	Name        string `"service" @Ident`
	Description string `  "description" @String`
	// Entries []*ServiceEntry `"{" @@* "}"`
}

type ServiceEntry struct {
	Pos lexer.Position

	Description *Description `( @@`
	Protocol    *Protocol    ` | @@`
	Registry    *Registry    ` | @@`
	Server      *MServer     ` | @@`
	Transport   *Transport   ` | @@ ) ";"*`
}

type Description struct {
	Pos  lexer.Position
	Name string `  "description" @Ident`
}
type Protocol struct {
	Pos  lexer.Position
	Name string `  "protocol" @Ident`
}
type Registry struct {
	Pos  lexer.Position
	Name string `  "registry" @Ident`
}
type MServer struct {
	Pos  lexer.Position
	Name string `  "server" @Ident`
}
type Transport struct {
	Pos  lexer.Position
	Name string `  "transport" @Ident`
}

type Endpoint struct {
	Pos lexer.Position

	Name  string           `"endpoint" @Ident`
	Entry []*EndpointEntry `"{" ( @@ ";"? )* "}"`
}

type EndpointEntry struct {
	Pos lexer.Position

	//	Option *Option `  "option" @@`
	Method *Method ` @@`
}

type Method struct {
	Pos lexer.Position

	Name              string `"rpc" @Ident`
	StreamingRequest  bool   `"(" @"stream"?`
	Request           *Type  `    @@ ")"`
	StreamingResponse bool   `"returns" "(" @"stream"?`
	Response          *Type  `              @@ ")"`
	// Options           []*Option `( "{" ( "option" @@ ";" )* "}" )?`
}

type Enum struct {
	Pos lexer.Position

	Name   string       `"enum" @Ident`
	Values []*EnumEntry `"{" ( @@ ( ";" )* )* "}"`
}

type EnumEntry struct {
	Pos lexer.Position

	Value  *EnumValue `  @@`
	Option *Option    `| "option" @@`
}

type EnumValue struct {
	Pos lexer.Position

	Key   string `@Ident`
	Value int    `"=" @( [ "-" ] Int )`

	Options []*Option `( "[" @@ ( "," @@ )* "]" )?`
}

type Message struct {
	Pos lexer.Position

	Name    string          `"type" @Ident`
	Entries []*MessageEntry `"{" @@* "}"`
}

type MessageEntry struct {
	Pos lexer.Position

	Enum    *Enum    `( @@`
	Message *Message ` | @@`
	Field   *Field   ` | @@ ) ";"*`
}

type Field struct {
	Pos lexer.Position

	Optional bool `(   @"optional"`
	Required bool `  | @"required"`
	Repeated bool `  | @"repeated" )?`

	Name string `@Ident`
	Type *Type  `@@`

	Options []*Option `( "[" @@ ( "," @@ )* "]" )?`
}

type Scalar int

const (
	None Scalar = iota
	Double
	Float
	Int
	Int32
	Int64
	Uint32
	Uint64
	Sint32
	Sint64
	Fixed32
	Fixed64
	SFixed32
	SFixed64
	Bool
	String
	Bytes
)

var scalarToString = map[Scalar]string{
	None: "None", Double: "Double", Float: "Float", Int: "int", Int32: "int32", Int64: "int64", Uint32: "uint32",
	Uint64: "uint64", Sint32: "Sint32", Sint64: "Sint64", Fixed32: "Fixed32", Fixed64: "Fixed64",
	SFixed32: "SFixed32", SFixed64: "SFixed64", Bool: "bool", String: "string", Bytes: "Bytes",
}

func (s Scalar) GoString() string { return scalarToString[s] }

var stringToScalar = map[string]Scalar{
	"double": Double, "float": Float, "int": Int, "int32": Int32, "int64": Int64, "uint32": Uint32, "uint64": Uint64,
	"sint32": Sint32, "sint64": Sint64, "fixed32": Fixed32, "fixed64": Fixed64, "sfixed32": SFixed32,
	"sfixed64": SFixed64, "bool": Bool, "string": String, "bytes": Bytes,
}

func (s *Scalar) Parse(lex *lexer.PeekingLexer) error {
	token := lex.Peek()
	v, ok := stringToScalar[token.Value]
	if !ok {
		return participle.NextMatch
	}
	lex.Next()
	*s = v
	return nil
}

type Type struct {
	Pos lexer.Position

	Scalar    Scalar   `  @@`
	Map       *MapType `| @@`
	Reference string   `| @(Ident ( "." Ident )*)`
}

type MapType struct {
	Pos lexer.Position

	Key   *Type `"map" "<" @@`
	Value *Type `"," @@ ">"`
}

var Parser = participle.MustBuild[Definition](participle.UseLookahead(2))

/*
func main() {
	ctx := kong.Parse(&cli)

	for _, file := range cli.Files {
		fmt.Println(file)
		r, err := os.Open(file)
		ctx.FatalIfErrorf(err, "")
		proto, err := parser.Parse("", r)
		ctx.FatalIfErrorf(err, "")
		repr.Println(proto, repr.Hide[lexer.Position]())
	}
}
*/
