package main 

import (
	"context"
	"fmt"
	"time"
)

// controlling timeouts
// cancelling goroutines
// passing metadata across your go app


func main(){
	ctx := context.Background()
	exampleTimeout(ctx)
	exampleWithValues()
}

func exampleTimeout(ctx context.Context){
	// ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func(){
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
		case <-done:
			fmt.Println("operation completed")
		case <-ctxWithTimeout.Done():
			fmt.Println("operation timed out", ctxWithTimeout.Err())
			// add logging or cleanup here
	}
}

func exampleWithValues() {
	type key string 
	const requestIDKey key = ""
	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, requestIDKey, "12345")

	if requestID, ok := ctxWithValue.Value(requestIDKey).(string); ok {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("Request ID not found")
	} 
}