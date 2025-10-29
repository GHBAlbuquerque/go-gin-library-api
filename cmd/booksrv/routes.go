package main

import (
	"example/go-gin-library-api/internal/book"

	"github.com/gin-gonic/gin"
)

func newRouter(h *book.BookHandler) (*gin.Engine, error) {
	router := gin.Default()

	router.GET("/books", h.FindAll)
	router.GET("/books/:id", h.GetById)
	router.POST("/books", h.Create)
	router.PATCH("/checkout", h.Checkout)
	router.PATCH("/return", h.Return)

	return router, nil
}
