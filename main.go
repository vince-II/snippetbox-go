package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// https://pkg.go.dev/net/http#pkg-constants

func home(w http.ResponseWriter, r *http.Request) {
	// check request URL if path only "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// programing
// /snippet/view?id=123
func snippetView(w http.ResponseWriter, r *http.Request) {
	//extract values id parameter from query string
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Display a specific snippet"))
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Set a new cache-control header. If an existing "Cache-Control" header exists
		// it will be overwritten.
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		// In contrast, the Add() method appends a new "Cache-Control" header and can
		// be called multiple times.
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")

		// Delete all values for the "Cache-Control" header.
		w.Header().Del("Cache-Control")

		// Retrieve the first value for the "Cache-Control" header.
		w.Header().Get("Cache-Control")

		// Retrieve a slice of all values for the "Cache-Control" header.
		w.Header().Values("Cache-Control")

		w.Header().Set("Allow", http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"name": "Alex"}`))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Createa a new snippet..."))
}

func main() {
	fmt.Println("Hello World")

	// routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

//curl -i -X POST http://localhost:4000/snippet/create
