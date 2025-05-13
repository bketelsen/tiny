// Package handlers contains the implementation of the <no value> service
package handlers

import (
	"encoding/json"
	"log"

	"_example"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

// SearchService is a struct for the SearchService endpoint
// It is the server implementation of the SearchServiceServer interface
// TODO: Add fields to the struct if needed for server dependencies and state
type SearchService struct {
	nc *nats.Conn
}

// Search is the implementation of the SearchService.Search endpoint
func (s *SearchService) Search(req micro.Request) {
	// Unmarshal the request
	input := &_example.SearchRequest{}
	err := json.Unmarshal(req.Data(), input)
	if err != nil {
		log.Println("Error unmarshalling request: ", err)
		return
	}

	// Create the response
	rsp := &_example.SearchResponse{}
	// TODO: implement the endpoint logic
	err = req.RespondJSON(rsp)
	if err != nil {
		log.Println("Error responding:", err)
		return
	}
	return
}

// NewSearchService creates a new SearchService struct
// TODO: Add parameters to the the function if needed to set server dependencies and state
func NewSearchService(nc *nats.Conn) *SearchService {
	return &SearchService{
		nc: nc,
	}
}
