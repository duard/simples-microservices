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
	log.Print("starting server...%s", dt)
	http.HandleFunc("/", handler)


	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("defaulting to port %s", port)
	}


	n := 7
	p := &n
	c := make(chan int)


	go func() {
		c <- *p + 0
	}()
	time.Sleep(time.Second)
	n = 8
	log.Print(<-c)

	go func() { c <- *p }()
	time.Sleep(time.Second)
	n = 9
	log.Print(<-c)

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
