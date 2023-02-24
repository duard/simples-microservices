package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/products/", app.all).Methods("GET")
	r.HandleFunc("/api/products/{id}", app.findByID).Methods("GET")
	r.HandleFunc("/api/products/", app.insert).Methods("POST")
	r.HandleFunc("/api/products/{id}", app.delete).Methods("DELETE")

	return r
}
