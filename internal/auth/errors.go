package auth

import "fmt"

var (
	ErrInvalidAuthType    = fmt.Errorf("unsupported authentication type")
	ErrInvalidRequest     = fmt.Errorf("missing required fields client_id and client_secret")
	ErrInvalidCredentials = fmt.Errorf("invalid credentials")
	ErrOnTokenIssue       = fmt.Errorf("could not issue token")
)
