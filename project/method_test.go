package project

import (
	"testing"

	"github.com/bketelsen/tiny/mucl"

	require "github.com/alecthomas/assert/v2"
)

func TestMethods(t *testing.T) {
	def, err := mucl.Parser.ParseString("", goodMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	searchService, ok := svc.GetEndpoint("SearchService")
	require.True(t, ok)
	require.Equal(t, "SearchService", searchService.Name)
	searchMethod, ok := searchService.GetMethod("Search")
	require.True(t, ok)
	require.Equal(t, "Search", searchMethod.Name)
	require.Equal(t, "SearchRequest", searchMethod.RequestTypeName)
	require.Equal(t, "SearchResponse", searchMethod.ResponseTypeName)
	require.Equal(t, false, searchMethod.RequestStreaming)
	require.Equal(t, false, searchMethod.ResponseStreaming)
	require.Equal(t, 1, len(svc.EndpointMap))
}

func TestStreamMethods(t *testing.T) {
	def, err := mucl.Parser.ParseString("", goodStreamMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	searchService, ok := svc.GetEndpoint("SearchService")
	require.True(t, ok)
	require.Equal(t, "SearchService", searchService.Name)
	searchMethod, ok := searchService.GetMethod("Search")
	require.True(t, ok)
	require.Equal(t, "Search", searchMethod.Name)
	require.Equal(t, "SearchRequest", searchMethod.RequestTypeName)
	require.Equal(t, "SearchResponse", searchMethod.ResponseTypeName)
	require.Equal(t, false, searchMethod.RequestStreaming)
	require.Equal(t, true, searchMethod.ResponseStreaming)
	require.Equal(t, 1, len(svc.EndpointMap))
}
