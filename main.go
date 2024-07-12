package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response represents a simple JSON response
type Response struct {
	Message string `json:"message"`
}

// helloHandler handles the /hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := ":443"
	certFile := "./certs/cert.pem"
	keyFile := "./certs/key.pem"

	log.Printf("Starting server on %s\n", port)

	http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServeTLS(port, certFile, keyFile, nil); err != nil {
		log.Fatal(err)
	}
}
