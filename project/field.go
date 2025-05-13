package project

import "github.com/iancoleman/strcase"

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

func (f *Field) DeclarationTag() string {
	return "`json:\"" + strcase.ToSnake(f.Name) + "\"`"
}
