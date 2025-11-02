package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Auth Service offers the necessary methods for Token creation and validation.
type Service interface {
	IssueToken(clientID string, ttl time.Duration) (TokenRes, error)
	ParseAndValidate(tokenStr string) (*Claims, error)
}

// JwtService is the implementation of AuthService by using Jwt.
type JwtService struct {
	secretKey []byte
	issuer    string
	audience  string
}

func NewService(secretKey, issuer, audience string) Service {
	j := JwtService{
		secretKey: []byte(secretKey),
		issuer:    issuer,
		audience:  audience,
	}

	return &j
}

// Issues a new token based on client credentials and desired ttl. Returns a TokenRes or an error.
func (s *JwtService) IssueToken(clientID string, ttl time.Duration) (TokenRes, error) {
	now := time.Now()
	exp := now.Add(ttl)
	claims := &Claims{
		ClientID: clientID,
		// still must write RegisteredClaims: jwt.RegisteredClaims{...} during initialization, because Go needs to know which anonymous field you’re populating.
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Audience:  []string{s.audience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := t.SignedString(s.secretKey)
	if err != nil {
		return TokenRes{}, err
	}

	return TokenRes{AccessToken: signed, ExpiresIn: int64(ttl.Seconds())}, nil
}

// ParseAndValidate parses a received token and returns its claims or an error.
func (s *JwtService) ParseAndValidate(tokenStr string) (*Claims, error) {
	//TODO

	return &Claims{}, nil
}
