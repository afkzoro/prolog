package main

import (
	"log"

	"github.com/afkzoro/prolog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":4000")
	log.Fatal(srv.ListenAndServe())
}