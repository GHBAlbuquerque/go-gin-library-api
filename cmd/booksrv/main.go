package main

import (
	"example/go-gin-library-api/internal/book"
	"example/go-gin-library-api/internal/book/stores"

	"github.com/gin-gonic/gin"
)

var books = []book.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func main() {
	router := gin.Default()
	h := book.NewHandler(stores.NewMemory(books))

	router.GET("/books", h.FindAll)
	router.GET("/books/:id", h.GetById)
	router.POST("/books", h.Create)
	router.PATCH("/checkout", h.Checkout)
	router.PATCH("/return", h.Return)

	router.Run("localhost:8080")

}
