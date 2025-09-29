package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// stack trace + writes error logs
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// custom error handler + human friendly error message
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// convenience wrapper around client error
func (app *application) notFound(w http.ResponseWriter, status int) {
	app.clientError(w, http.StatusNotFound)
}
