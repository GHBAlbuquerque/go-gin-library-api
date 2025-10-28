package main

import (
	"example/go-gin-library-api/internal/book"
	"example/go-gin-library-api/internal/bootstrap"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	s, err := bootstrap.NewStoreFromEnv()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	h := book.NewHandler(s)

	router.GET("/books", h.FindAll)
	router.GET("/books/:id", h.GetById)
	router.POST("/books", h.Create)
	router.PATCH("/checkout", h.Checkout)
	router.PATCH("/return", h.Return)

	router.Run("localhost:8080")
}

//TODO: add dynamic id generation for book creation using UUID
