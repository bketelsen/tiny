package _example

// GetRequest is a struct for the GetRequest type
type GetRequest struct {
	UID int `json:"uid" yaml:"uid"`
}

// GetResponse is a struct for the GetResponse type
type GetResponse struct {
	FirstName string `json:"first_name" yaml:"first_name"`
	LastName  string `json:"last_name" yaml:"last_name"`
}

// UnlockRequest is a struct for the UnlockRequest type
type UnlockRequest struct {
	UID int `json:"uid" yaml:"uid"`
}

// UnlockResponse is a struct for the UnlockResponse type
type UnlockResponse struct{}
