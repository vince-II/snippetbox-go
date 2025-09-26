package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// go run ./cmd/web -addr=":9999"
	addr := flag.String("addr", ":4000", "HTTP network address")
	addr := os.Getenv("SNIPPETBOX_ADDR")
	//
	flag.Parse()
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
	mux.HandleFunc("/snippet/download", downloadHandler)

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)

}
