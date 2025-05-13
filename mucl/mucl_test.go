package mucl

import (
	"testing"

	require "github.com/alecthomas/assert/v2"
)

func TestGood(t *testing.T) {
	tree, err := Parser.ParseString("", goodMucl)
	require.NoError(t, err)
	require.Equal(t, 6, len(tree.Entries))
	require.Equal(t, "helloworld", tree.ServiceName())

	require.Equal(t, "something", tree.Description())
	require.Equal(t, "SearchRequest", tree.Entries[2].Message.Name)
	require.Equal(t, "SearchResponse", tree.Entries[3].Message.Name)
	require.Equal(t, "SearchType", tree.Entries[4].Enum.Name)
	require.Equal(t, "SearchService", tree.Entries[5].Endpoint.Name)
}

func TestEmbeddedGood(t *testing.T) {
	tree, err := Parser.ParseString("", embeddedEnumMucl)
	require.NoError(t, err)
	require.Equal(t, 5, len(tree.Entries))
	require.Equal(t, "helloworld", tree.ServiceName())

	require.Equal(t, "something", tree.Description())
	require.Equal(t, "SearchRequest", tree.Entries[2].Message.Name)
	require.Equal(t, "SearchResponse", tree.Entries[3].Message.Name)
	require.Equal(t, "SearchService", tree.Entries[4].Endpoint.Name)
}

func TestGoodTwoEndpoints(t *testing.T) {
	tree, err := Parser.ParseString("", twoEndpoints)
	require.NoError(t, err)
	require.Equal(t, 7, len(tree.Entries))
	require.Equal(t, "helloworld", tree.ServiceName())

	require.Equal(t, "something", tree.Description())
	require.Equal(t, "SearchRequest", tree.Entries[2].Message.Name)
	require.Equal(t, "SearchResponse", tree.Entries[3].Message.Name)
	require.Equal(t, "SearchType", tree.Entries[4].Enum.Name)

	require.Equal(t, "SearchService", tree.Entries[5].Endpoint.Name)
	require.Equal(t, "InternalSearchService", tree.Entries[6].Endpoint.Name)
}

func TestWithTwoServices(t *testing.T) {
	_, err := Parser.ParseString("", twoServices)
	require.NoError(t, err)
}

func TestWithOptions(t *testing.T) {
	tree, err := Parser.ParseString("", optionMucl)
	require.NoError(t, err)
	require.Equal(t, 6, len(tree.Entries))
	require.Equal(t, "helloworld", tree.ServiceName())
	require.Equal(t, "SearchRequest", tree.Entries[2].Message.Name)
	require.Equal(t, "SearchResponse", tree.Entries[3].Message.Name)
	require.Equal(t, "SearchType", tree.Entries[4].Enum.Name)
	require.Equal(t, "SearchService", tree.Entries[5].Endpoint.Name)
}

var goodMucl = `
service="helloworld"
description="something"

type SearchRequest {
  query string
  type SearchType
  page_number int32
  result_per_page int32
  intpage int
}

type SearchResponse {
  results string
}

enum SearchType {
  SHALLOW = 0
  DEEP = 1
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`

var twoServices = `
service="helloworld"
description="something"
service="helloworld"
description="something"

type SearchRequest {
  query string
  type SearchType
  page_number int32
  result_per_page int32
}

type SearchResponse {
  results string
}

enum SearchType {
  SHALLOW = 0
  DEEP = 1
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`

var twoEndpoints = `
service="helloworld"
description="something"

type SearchRequest {
  query string
  type SearchType
  page_number int32
  result_per_page int32
}

type SearchResponse {
  results string
}

enum SearchType {
  SHALLOW = 0
  DEEP = 1
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}

endpoint InternalSearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`

var embeddedEnumMucl = `
service="helloworld"
description="something"

type SearchRequest {
  query string
  enum SearchType {
    SHALLOW = 0
    DEEP = 1
  }
  page_number int32
  result_per_page int32
}

type SearchResponse {
  results string
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`

var optionMucl = `
service="helloworld"
description="something"

type SearchRequest {
  query string
  type SearchType
  page_number int32
  result_per_page int32
}

type SearchResponse {
  results string
}

enum SearchType {
  SHALLOW = 0
  DEEP = 1
}

endpoint SearchService {
  rpc Search(SearchRequest) returns (SearchResponse)
}
`
