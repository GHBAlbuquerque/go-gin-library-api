package auth

import "github.com/golang-jwt/jwt/v5"

type AuthReq struct {
	ClientID     string `json:"client_id" binding:required`
	ClientSecret string `json:"client_secret" binding: required`
}

type TokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type Claims struct {
	ClientID string `json:"cid"`
	Claims   jwt.RegisteredClaims
}
