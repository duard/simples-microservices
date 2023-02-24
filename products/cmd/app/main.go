package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/duard/simples-microservices/products/pkg/models/mongodb"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	product *mongodb.ProductModel
}
func main() {
	dt := time.Now()
	fmt.Println("Current date and time is: ", dt.String())
	log.Print("starting server...")
	http.HandleFunc("/", handler)
	myMap := map[string]string{"first key": "first value", "second key": "second value", "third key": "third value", "fourth key": "fourth value", "fifth key": "fifth value"}
	log.Print(myMap)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server...
	log.Printf("Products listening on port http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, ":: Products %s!\n", name)
}
