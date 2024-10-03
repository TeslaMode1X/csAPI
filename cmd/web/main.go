package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":4000",
		Handler: router(),
	}

	log.Fatal(srv.ListenAndServe())
}
