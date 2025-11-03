package main

import (
	"example/go-gin-library-api/internal/auth"
	"example/go-gin-library-api/internal/book"

	"github.com/gin-gonic/gin"
)

func newRouter(authHandler *auth.Handler, bookHandler *book.Handler) (*gin.Engine, error) {
	router := gin.Default()

	router.POST("/auth/token", authHandler.RequestAuth)

	api := router.Group("/api", authHandler.RequireAuth())
	{
		api.GET("/books", bookHandler.FindAll)
		api.GET("/books/:id", bookHandler.GetById)
		api.POST("/books", bookHandler.Create)
		api.PATCH("/checkout", bookHandler.Checkout)
		api.PATCH("/return", bookHandler.Return)
	}

	return router, nil
}
