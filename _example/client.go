// Package _example defines the types and interfaces for the users service
package _example

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

// User Methods

func UserGet(nc *nats.Conn, in GetRequest) (GetResponse, error) {
	bb, err := json.Marshal(in)
	if err != nil {
		return GetResponse{}, err
	}
	var out GetResponse
	resp, err := nc.Request("users.UserGet", bb, nats.DefaultTimeout)
	if err != nil {
		return GetResponse{}, err
	}
	err = json.Unmarshal(resp.Data, &out)
	if err != nil {
		return GetResponse{}, err
	}
	return out, nil
}

func UserUnlock(nc *nats.Conn, in UnlockRequest) (UnlockResponse, error) {
	bb, err := json.Marshal(in)
	if err != nil {
		return UnlockResponse{}, err
	}
	var out UnlockResponse
	resp, err := nc.Request("users.UserUnlock", bb, nats.DefaultTimeout)
	if err != nil {
		return UnlockResponse{}, err
	}
	err = json.Unmarshal(resp.Data, &out)
	if err != nil {
		return UnlockResponse{}, err
	}
	return out, nil
}
