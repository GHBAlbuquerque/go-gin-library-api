package auth

import "github.com/golang-jwt/jwt/v5"

type TokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type Claims struct {
	ClientID             string `json:"cid"`
	jwt.RegisteredClaims        // embedded field of RegisteredClaims inside my struct
}

/*
	{
	"cid": "frontend",
	"exp": 1730490000,
	"iss": "go-gin-library-api"
	}
*/
