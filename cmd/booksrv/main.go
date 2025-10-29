package main

import (
	"example/go-gin-library-api/internal/book"
	"example/go-gin-library-api/internal/bootstrap"
	"log"
)

func main() {

	s, err := bootstrap.NewStoreFromEnv()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	sv := book.NewService(s)
	h := book.NewHandler(sv)
	r, err := newRouter(h)
	if err != nil {
		log.Fatal(err.Error())
	}

	r.Run("localhost:8080")
}

//TODO: add dynamic id generation for book creation using UUID
