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

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
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

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	// log.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
