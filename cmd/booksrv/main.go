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

	h := book.NewHandler(s)
	r, err := newRouter(h)
	if err != nil {
		log.Fatal(err.Error())
	}

	r.Run("localhost:8080")
}

//TODO: add dynamic id generation for book creation using UUID
// TODO: refactor to put logic on service package
