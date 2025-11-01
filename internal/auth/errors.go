package auth

import "fmt"

var (
	ErrInvalidRequest = fmt.Errorf("missing required fields client_id and client_secret")
)
