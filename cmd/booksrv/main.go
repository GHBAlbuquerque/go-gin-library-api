package main

import (
	"example/go-gin-library-api/internal/bootstrap"
	"log"
)

func main() {
	config, err := bootstrap.LoadAuthConfigFromEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	authHandler, bookHandler, err := bootstrap.BuildDeps(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	r, err := newRouter(authHandler, bookHandler)
	if err != nil {
		log.Fatal(err.Error())
	}

	r.Run(config.Addr)
}
