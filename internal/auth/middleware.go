package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireAuth ensures requests carry a valid bearer token before reaching the
// protected handlers.
func (h *Handler) RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		claims, err := h.service.ParseAndValidate(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("client_id", claims.ClientID)
		ctx.Next()
	}
}
