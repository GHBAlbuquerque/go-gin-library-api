package book

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	store Store // it's illegal to use pointers to interface types
}

func NewHandler(store Store) *BookHandler {
	h := BookHandler{
		store: store,
	}

	return &h
}

/*getBooks returns the json version of my book slice*/
func (h *BookHandler) FindAll(c *gin.Context) {
	out := h.store.List(c)
	c.IndentedJSON(http.StatusOK, out)
}

/*getBookById returns the json version of desired book */
func (h *BookHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	book, err := h.store.FindById(c, id)

	if err != nil {
		text := fmt.Sprintf("Book with id %s not found", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": text})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

/*createBook creates a book and the json version of my book slice*/
func (h *BookHandler) Create(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		text := fmt.Sprintf("BindJSON: %s", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": text})
		return
	}

	out, err := h.store.Create(c, newBook)
	if err != nil {
		text := fmt.Sprintf("Create: %s", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": text})
		return
	}

	c.IndentedJSON(http.StatusCreated, out)
}

/*checkoutBook retrieves an available book from the library*/
func (h *BookHandler) Checkout(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No id sent"})
		return
	}

	book, err := h.store.FindById(c, id)

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
	h.store.Update(c, book)
	c.IndentedJSON(http.StatusOK, book)
}

/*returnBook retrieves an available book from the library*/
func (h *BookHandler) Return(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No id sent"})
		return
	}

	book, err := h.store.FindById(c, id)

	if err != nil {
		text := fmt.Sprintf("Book with id %s not found", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": text})
		return
	}

	book.Quantity += 1
	h.store.Update(c, book)
	c.IndentedJSON(http.StatusOK, book)
}
