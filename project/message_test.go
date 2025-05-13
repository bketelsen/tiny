package project

import (
	"testing"

	"github.com/bketelsen/tiny/mucl"

	require "github.com/alecthomas/assert/v2"
)

func TestMessages(t *testing.T) {
	def, err := mucl.Parser.ParseString("", goodMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 2, len(svc.MessageMap))
}

func TestEmbeddedEnumMessages(t *testing.T) {
	def, err := mucl.Parser.ParseString("", embeddedEnumMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 2, len(svc.MessageMap))
	searchType, ok := svc.GetMessage("SearchRequest")
	require.True(t, ok)
	require.Equal(t, "SearchRequest", searchType.Name)
	require.Equal(t, 4, len(searchType.GetAllFields()))
}

func TestEmbeddedMsgMessages(t *testing.T) {
	def, err := mucl.Parser.ParseString("", embeddedTypeMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 3, len(svc.MessageMap))
	responseType, ok := svc.GetMessage("SearchResponse")
	require.True(t, ok)
	require.Equal(t, "SearchResponse", responseType.Name)
	require.Equal(t, 2, len(responseType.GetAllFields()))
}
