package main 

import (
	"context"
	"fmt"
	"time"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(),4*time.Second)
	defer cancel()

	select {
	case <- time.After(3 * time.Second):
		fmt.Println("Finished processing request")
		fmt.Fprintln(w, "Hello, world!")
	case <- ctx.Done():
		fmt.Println("Request timed out")
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		return
	}
}