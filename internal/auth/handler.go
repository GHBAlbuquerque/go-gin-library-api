package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	a := AuthHandler{}

	return &a
}

func (a *AuthHandler) RequestAuth(ctx *gin.Context) {
	clientId := ctx.Query("client_id")
	clientSecret := ctx.Query("client_secret")

	if clientId == "" || clientSecret == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidRequest.Error()})
		return
	}

	authReq := AuthReq{ClientID: clientId, ClientSecret: clientSecret}

	//TODO
	ctx.IndentedJSON(http.StatusOK, authReq)
}
