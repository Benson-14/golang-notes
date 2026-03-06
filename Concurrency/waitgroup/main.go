package main

import (
	"fmt"
	"sync"
)

// waitgroup --> ADD, DONE, WAIT
// so for per goroutine we use use ADD and in the goroutine func, we use Done to decrement the counter
// and in the main func, we use Wait to wait for all the goroutines to finish

// The three operations:

// Add(n) - Increment the counter by n (call before launching goroutines)
// Done() - Decrement the counter by 1 (call when goroutine finishes)
// Wait() - Block until the counter reaches 0

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when done
	fmt.Printf("Worker %d completed\n", id)
}

func go_worker(id int) {
	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1) // Increment counter
		go worker(i+1, &wg)
	}

	wg.Wait() // Block until counter is 0
	fmt.Println("All workers finished!")

	// Go 1.25 introduced a simpler way to launch goroutines with WaitGroup.
	// The new Go() method combines Add(1), launching the goroutine,
	// and defer Done() into a single call:

	// for i := range 3 {
	// 	wg.Go(func() {
	// 		fmt.Printf("Worker %d done\n", i+1)
	// 	})
	// }

	for i := range 5 {
		wg.Go(func() {
			go_worker(i + 1)
		})
	}

	wg.Wait()
	fmt.Println("All workers finished!")
}
