package project

import (
	"testing"

	"github.com/bketelsen/tiny/mucl"

	require "github.com/alecthomas/assert/v2"
)

func TestEndpoint(t *testing.T) {
	def, err := mucl.Parser.ParseString("", goodMucl)
	require.NoError(t, err)
	svc, err := fromMuCL(def)
	require.NoError(t, err)
	require.Equal(t, 1, len(svc.EndpointMap))
	searchService, ok := svc.GetEndpoint("SearchService")
	require.True(t, ok)
	require.Equal(t, "SearchService", searchService.Name)
}
