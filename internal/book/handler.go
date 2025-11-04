package book

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	h := Handler{
		service: service,
	}

	return &h
}

// FindAll returns the json version of my book slice.
func (h *Handler) FindAll(ctx *gin.Context) {
	if ctx.Query("title") != "" && ctx.Query("author") != "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidFilter.Error()})
		return
	}

	filters := BookFilters{Author: ctx.Query("author"), Title: ctx.Query("title")}

	books, err := h.service.FindAll(ctx.Request.Context(), filters)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	out := make([]BookResponse, 0, len(books))
	for _, book := range books {
		out = append(out, BookResponse(book))
	}

	ctx.IndentedJSON(http.StatusOK, out)
}

// GetById returns the json version of desired book.
func (h *Handler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := h.service.GetById(ctx.Request.Context(), id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, BookResponse(book))
}

// Create creates a book and the json version of my book slice.
func (h *Handler) Create(ctx *gin.Context) {
	var bookRequest BookRequest

	if err := ctx.BindJSON(&bookRequest); err != nil {
		text := fmt.Sprintf("BindJSON: %s", err.Error())
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": text})
		return
	}

	out, err := h.service.Create(ctx.Request.Context(), bookRequest)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, out)
}

// Checkout retrieves an available book from the library.
func (h *Handler) Checkout(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No id sent"})
		return
	}

	book, err := h.service.Checkout(ctx.Request.Context(), id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, BookResponse(book))
}

// Return retrieves an available book from the library.
func (h *Handler) Return(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No id sent"})
		return
	}

	book, err := h.service.Return(ctx.Request.Context(), id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, BookResponse(book))
}
