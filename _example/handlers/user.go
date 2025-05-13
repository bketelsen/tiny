// Package handlers contains the implementation of the <no value> service
package handlers

import (
	"encoding/json"
	"log"

	"_example"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

// User is a struct for the User endpoint
// It is the server implementation of the UserServer interface
// TODO: Add fields to the struct if needed for server dependencies and state
type User struct {
	nc *nats.Conn
}

// Get is the implementation of the User.Get endpoint
func (s *User) Get(req micro.Request) {
	// Unmarshal the request
	input := &_example.GetRequest{}
	err := json.Unmarshal(req.Data(), input)
	if err != nil {
		log.Println("Error unmarshalling request: ", err)
		return
	}

	// Create the response
	rsp := &_example.GetResponse{}
	// TODO: implement the endpoint logic
	err = req.RespondJSON(rsp)
	if err != nil {
		log.Println("Error responding:", err)
		return
	}
	return
}

// Unlock is the implementation of the User.Unlock endpoint
func (s *User) Unlock(req micro.Request) {
	// Unmarshal the request
	input := &_example.UnlockRequest{}
	err := json.Unmarshal(req.Data(), input)
	if err != nil {
		log.Println("Error unmarshalling request: ", err)
		return
	}

	// Create the response
	rsp := &_example.UnlockResponse{}
	// TODO: implement the endpoint logic
	err = req.RespondJSON(rsp)
	if err != nil {
		log.Println("Error responding:", err)
		return
	}
	return
}

// NewUser creates a new User struct
// TODO: Add parameters to the the function if needed to set server dependencies and state
func NewUser(nc *nats.Conn) *User {
	return &User{
		nc: nc,
	}
}
