package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("listening server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
