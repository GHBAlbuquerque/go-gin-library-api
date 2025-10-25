package main

import (
	"example/go-gin-library-api/internal/book"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", book.FindAll)
	router.GET("/books/:id", book.GetById)
	router.POST("/books", book.Create)
	router.PATCH("/checkout", book.Checkout)
	router.PATCH("/return", book.Return)

	router.Run("localhost:8080")

}
