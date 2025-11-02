package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository ClientRepository
	service    Service
}

// NewHandler creates the Auth endpoint for authentication request.
func NewHandler(repository ClientRepository, service Service) *Handler {
	a := Handler{
		repository: repository,
		service:    service,
	}

	return &a
}

// RequestAuth returns a Bearer Token if credentials received on client_id and client_secret are valid.
func (h *Handler) RequestAuth(ctx *gin.Context) {

	grantType := ctx.PostForm("grant_type")
	if grantType != "client_credentials" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidAuthType.Error()})
		return
	}

	clientID := ctx.PostForm("client_id")
	clientSecret := ctx.PostForm("client_secret")

	if clientID == "" || clientSecret == "" {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidRequest.Error()})
		return
	}

	if exists := h.repository.Validate(clientID, clientSecret); !exists {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidCredentials.Error()})
		return
	}

	tok, err := h.service.IssueToken(clientID, time.Second)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": ErrOnTokenIssue.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, tok)
}
