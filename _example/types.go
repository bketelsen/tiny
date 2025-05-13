package _example

// SearchRequest is a struct for the SearchRequest type
type SearchRequest struct { 
  PageNumber int `json:"page_number"`
  Query string `json:"query"`
  ResultPerPage int `json:"result_per_page"`
  Type SearchType `json:"type"`
}
// SearchResponse is a struct for the SearchResponse type
type SearchResponse struct { 
  Results string `json:"results"`
}

// SearchType is a type for the SearchType enum
type SearchType int

const (
	SHALLOW SearchType = 0
	DEEP SearchType = 1
)

