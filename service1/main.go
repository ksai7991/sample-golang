package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofrs/uuid"
)

func logRequest(r *http.Request) {
	uri := r.RequestURI
	method := r.Method
	fmt.Println("Got request!", method, uri)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		fmt.Fprintf(w, "Hello from Service 1! You've requested %s\n", r.URL.Path)
	})

	http.HandleFunc("/uuid", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		requestID := uuid.Must(uuid.NewV4())
		fmt.Fprint(w, requestID.String())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	fmt.Printf("Service 1 is running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
