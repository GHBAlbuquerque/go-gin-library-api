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

/*FindAll returns the json version of my book slice*/
func (h *BookHandler) FindAll(ctx *gin.Context) {
	if ctx.Query("title") != "" && ctx.Query("author") != "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Either filter by title OR author"})
		return
	}

	if title := ctx.Query("title"); title != "" {
		books, err := h.store.FindByTitle(ctx, title)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, books)
		return
	}

	if author := ctx.Query("author"); author != "" {
		books, err := h.store.FindByAuthor(ctx, author)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, books)
		return
	}

	out, err := h.store.List(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, out)
}

/*GetById returns the json version of desired book */
func (h *BookHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := h.store.FindById(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, book)
}

/*Create creates a book and the json version of my book slice*/
func (h *BookHandler) Create(ctx *gin.Context) {
	var newBook Book

	if err := ctx.BindJSON(&newBook); err != nil {
		text := fmt.Sprintf("BindJSON: %s", err.Error())
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": text})
		return
	}

	out, err := h.store.Create(ctx, newBook)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, out)
}

/*Checkout retrieves an available book from the library*/
func (h *BookHandler) Checkout(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No id sent"})
		return
	}

	book, err := h.store.FindById(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if book.Quantity == 0 {
		text := fmt.Sprintf("Book '%s' is not available", book.Title)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": text})
		return
	}

	book.Quantity -= 1
	if err := h.store.Update(ctx, book); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, book)
}

/*Return retrieves an available book from the library*/
func (h *BookHandler) Return(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No id sent"})
		return
	}

	book, err := h.store.FindById(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	book.Quantity += 1
	if err := h.store.Update(ctx, book); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, book)
}
