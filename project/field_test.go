package project

import (
	"slices"
	"testing"

	"github.com/bketelsen/tiny/mucl"

	require "github.com/alecthomas/assert/v2"
)

func TestFields(t *testing.T) {
	def, err := mucl.Parser.ParseString("", goodMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 2, len(svc.MessageMap))
	searchType, ok := svc.GetMessage("SearchRequest")
	require.True(t, ok)
	require.Equal(t, "SearchRequest", searchType.Name)
	allFields := searchType.GetAllFields()
	require.Equal(t, 4, len(allFields))
	fieldNames := searchType.GetFieldNames()
	contains := slices.Contains(fieldNames, "type")
	require.True(t, contains)
	fieldTypes := searchType.GetFieldTypes()
	contains = slices.Contains(fieldTypes, "SearchType")
	require.True(t, contains)
	contains = slices.Contains(fieldTypes, "string")
	require.True(t, contains)
	contains = slices.Contains(fieldTypes, "int32")
	require.True(t, contains)
	contains = slices.Contains(fieldTypes, "bool")
	require.False(t, contains)
}
