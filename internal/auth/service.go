package auth

import "time"

// Auth Service offers the necessary methods for Token creation and validation.
type Service interface {
	IssueToken(clientID, clientSecret string, ttl time.Duration) (TokenRes, error)
	ParseAndValidate(tokenStr string) (*Claims, error)
}

// JwtService is the implementation of AuthService by using Jwt.
type JwtService struct {
	secret   []byte
	issuer   string
	audience string
}

func NewService(secret, issuer, audience string) Service {
	j := JwtService{
		secret:   []byte(secret),
		issuer:   issuer,
		audience: audience,
	}

	return &j
}

// Issues a new token based on client credentials and desired ttl. Returns a TokenRes or an error.
func (j *JwtService) IssueToken(clientID, clientSecret string, ttl time.Duration) (TokenRes, error) {
	//TODO

	return TokenRes{}, nil
}

// ParseAndValidate parses a received token and returns its claims or an error.
func (j *JwtService) ParseAndValidate(tokenStr string) (*Claims, error) {
	//TODO

	return &Claims{}, nil
}
