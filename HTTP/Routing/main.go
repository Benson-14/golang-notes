package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Path: /users/123
	parts := strings.Split(r.URL.Path, "/")
	// parts = ["", "users", "123"]

	if len(parts) < 3 {
		http.Error(w, "Missing userID", http.StatusBadRequest)
		return
	}

	// id := parts[2]
	id := r.PathValue("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User id: %d", userID)

}

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /users/{id}", getUserHandler)
	mux.HandleFunc("GET /users/{id}", getUserHandler)
	// mux.HandleFunc("POST /users/", createUserHandler)

	fmt.Println("Server has started....")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
