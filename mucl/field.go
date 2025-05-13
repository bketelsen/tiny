package mucl

import "github.com/iancoleman/strcase"

func (f *Field) ExportedName() string {
	if f == nil {
		return ""
	}
	if f.Name == "" {
		return ""
	}
	return strcase.ToCamel(f.Name)
}
