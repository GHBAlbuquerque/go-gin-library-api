package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

/*getBooks returns the json version of my book slice*/
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

/*createBook creates a book and the json version of my book slice*/
func createBook(c *gin.Context) {
	var newBook book

	err := c.BindJSON(&newBook)

	if err != nil {
		return
	}

	for _, book := range books {
		if book.ID == newBook.ID {
			c.JSON(http.StatusConflict, gin.H{"error": "Id already exists"})
			return
		}

		if book.Title == newBook.Title && book.Author == newBook.Author {
			c.JSON(http.StatusConflict, gin.H{"error": "Title from this Author already exists"})
			return
		}
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", createBook)

	router.Run("localhost:8080")

}
