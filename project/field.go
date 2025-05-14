package project

import (
	"strings"

	"github.com/iancoleman/strcase"
)

type Field struct {
	Optional bool
	Repeated bool
	Required bool
	Name     string
	TypeName string
}

func init() {
	strcase.ConfigureAcronym("api", "API")
	strcase.ConfigureAcronym("uid", "UID")
	strcase.ConfigureAcronym("gid", "GID")
	strcase.ConfigureAcronym("id", "ID")
	strcase.ConfigureAcronym("uuid", "UUID")
	strcase.ConfigureAcronym("db", "DB")
}

func (f *Field) DeclarationType() string {
	if f.Repeated {
		return "[]" + f.TypeName
	}
	if f.Required {
		return "*" + f.TypeName
	}
	return f.TypeName
}

func (f *Field) DeclarationName() string {
	return strcase.ToCamel(f.Name)
}

func (f *Field) DeclarationTag(strct string) string {
	var sb strings.Builder
	sb.WriteString("`json:\"")
	sb.WriteString(strcase.ToSnake(f.Name))
	sb.WriteString("\" yaml:\"")
	sb.WriteString(strcase.ToSnake(f.Name))
	sb.WriteString("\"")
	if strct != "" {
		sb.WriteString(" env:\"")
		sb.WriteString(strcase.ToScreamingSnake(strct + " " + f.Name))
		sb.WriteString("\"")
		sb.WriteString(" env-description:\"")
		sb.WriteString(strct + " " + f.Name)
		sb.WriteString("\"")
	}
	sb.WriteString("`")
	return sb.String()
}
