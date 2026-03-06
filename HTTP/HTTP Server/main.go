package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	book := Book{Title: "The Go Programming Language", Author: "Donovan & Kernighan", Price: 39.99}
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK) not required default to 200 when you call encode
	json.NewEncoder(w).Encode(book)
	// json.NewEncoder(w) --> creates a new *json.Encoder that writes to w
	// .Encode(book) --> serializes `book` to JSON → writes it to w, returns error

}

// type ResponseWriter interface {
//     Header() http.Header        // Get headers map to modify
//     Write([]byte) (int, error)  // Write body bytes
//     WriteHeader(statusCode int) // Set status code
// }

// w.Header().Set("Content-Type", "application/json")  // Replace any existing value
// w.Header().Add("Cache-Control", "no-cache")         // Add to existing values

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my server")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/book", BookHandler)

	fmt.Println("Server has started....")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
