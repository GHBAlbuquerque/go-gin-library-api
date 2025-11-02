package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

// NewHandler creates the Auth endpoint for authentication request.
func NewHandler() *Handler {
	a := Handler{}

	return &a
}

// RequestAuth returns a Bearer Token if credentials received on client_id and client_secret are valid.
func (a *Handler) RequestAuth(ctx *gin.Context) {
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
