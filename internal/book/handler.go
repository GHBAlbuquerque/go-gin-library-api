package book

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service Service
}

func NewHandler(service Service) *BookHandler {
	h := BookHandler{
		service: service,
	}

	return &h
}

/*FindAll returns the json version of my book slice*/
func (h *BookHandler) FindAll(ctx *gin.Context) {
	if ctx.Query("title") != "" && ctx.Query("author") != "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidFilter.Error()})
		return
	}

	filters := BookFilters{Author: ctx.Query("author"), Title: ctx.Query("title")}

	out, err := h.service.FindAll(ctx, filters)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, out)
}

/*GetById returns the json version of desired book */
func (h *BookHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := h.service.GetById(ctx, id)

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

	out, err := h.service.Create(ctx, newBook)
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

	book, err := h.service.Checkout(ctx, id)
	if err != nil {
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

	book, err := h.service.Return(ctx, id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, book)
}
