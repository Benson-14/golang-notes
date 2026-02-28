package main

import (
	"fmt"
	"sync"
)

// mu.Lock()    // Acquire the lock (blocks if another goroutine holds it)
// counter++    // Only one goroutine can be here at a time
// mu.Unlock()  // Release the lock (allows another goroutine to acquire it)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	for range 1000 {
		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
		})
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}
