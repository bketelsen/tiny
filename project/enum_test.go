package project

import (
	"testing"

	"github.com/bketelsen/tiny/mucl"

	require "github.com/alecthomas/assert/v2"
)

func TestEnums(t *testing.T) {
	def, err := mucl.Parser.ParseString("", goodMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 1, len(svc.EnumMap))
	searchType, ok := svc.GetEnum("SearchType")
	require.True(t, ok)
	require.Equal(t, "SearchType", searchType.Name)
	require.Equal(t, 2, len(searchType.Values))
	require.Equal(t, "SHALLOW", searchType.Values[0].Key)
}
