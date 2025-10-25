package book

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	store Store // its illegal to use pointers to interface types
}

func NewHandler(store Store) *BookHandler {
	h := BookHandler{
		store: store,
	}

	return &h
}

/*getBooks returns the json version of my book slice*/
func (h *BookHandler) FindAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

/*getBookById returns the json version of desired book */
func (h *BookHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	book, err := FindbyId(id)

	if err != nil {
		text := fmt.Sprintf("Book with id %s not found", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": text})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

/*aux method*/
func (h *BookHandler) FindbyId(id string) (*Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

/*createBook creates a book and the json version of my book slice*/
func (h *BookHandler) Create(c *gin.Context) {
	var newBook Book

	err := c.BindJSON(&newBook)

	if err != nil {
		return
	}

	for _, book := range books {
		if book.ID == newBook.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "Id already exists"})
			return
		}

		if book.Title == newBook.Title && book.Author == newBook.Author {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "Title from this Author already exists"})
			return
		}
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

/*checkoutBook retrieves an available book from the library*/
func (h *BookHandler) Checkout(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No id sent"})
		return
	}

	book, err := FindbyId(id)

	if err != nil {
		text := fmt.Sprintf("Book with id %s not found", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": text})
		return
	}

	if book.Quantity == 0 {
		text := fmt.Sprintf("Book '%s' is not available", book.Title)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": text})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

/*returnBook retrieves an available book from the library*/
func (h *BookHandler) Return(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No id sent"})
		return
	}

	book, err := FindbyId(id)

	if err != nil {
		text := fmt.Sprintf("Book with id %s not found", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": text})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}
