package main

import (
	"encoding/json"
	"net/http"

	"github.com/duard/simples-microservices/products/pkg/models"
	"github.com/gorilla/mux"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	// Get all product stored
	products, err := app.product.All()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert product list into json encoding
	b, err := json.Marshal(products)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Products have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find product by id
	m, err := app.product.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Product not found")
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
	}

	// Convert product to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a product")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	// Define product model
	var m models.Product
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new product
	insertResult, err := app.product.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New product have been created, id=%s", insertResult.InsertedID)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete product by id
	deleteResult, err := app.product.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d product(s)", deleteResult.DeletedCount)
}
