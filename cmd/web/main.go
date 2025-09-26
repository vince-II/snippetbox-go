package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// file server that serves static
	fs := http.FileServer(http.Dir("./ui/static/"))

	// mux.Handle -> register file server as handler for all URL paths with
	// "/static/" . Note: if matched, remove the prefix before the request
	// reaches the file server
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	// routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
