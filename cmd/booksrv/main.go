package main

import (
	"example/go-gin-library-api/internal/auth"
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
	a := auth.NewHandler()
	r, err := newRouter(a, h)
	if err != nil {
		log.Fatal(err.Error())
	}

	r.Run("localhost:8080")
}

//TODO: add authentication and cors on middleware layer
//TODO: unit testing
