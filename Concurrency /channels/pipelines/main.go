package main

import (
	"fmt"
	"sync"
)

func GenerateNumbers(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := range n {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func SumChannel(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()

	return out
}

// ... means it accepts any number of channels (variadic)
// So channels is a slice of channels: []<-chan int

func Merge(channels ...<-chan int) <-chan int {
	// Forward all values from all input channels to output
	out := make(chan int)
	var wg sync.WaitGroup // will track the goroutines

	for _, ch := range channels { // each ch gets it's own goroutine
		wg.Go(func() {
			for v := range ch {
				out <- v
			}
		})
	}

	// wg.Go() does three things automatically:

	// wg.Add(1) — increments the counter
	// go func(){}() — launches the goroutine
	// wg.Done() — decrements counter when the goroutine finishes

	// Close output when all inputs are exhausted
	go func() {
		wg.Wait()  // blocks until all goroutines finish
		close(out) // close the output channel // merged will be closed here
	}()

	return out // 0, 1, 2, 0, 1, 2 (in any order)

}

func main() {
	// Pipeline: Generate -> Square -> Sum
	result := <-SumChannel(Square(GenerateNumbers(5)))
	fmt.Printf("Sum of squares 0-4: %d\n", result) // 0+1+4+9+16 = 30

	// Merge multiple generators
	merged := Merge(GenerateNumbers(3), GenerateNumbers(3))
	sum := 0
	for n := range merged {
		sum += n
	}
	fmt.Printf("Merged sum: %d\n", sum) // (0+1+2)*2 = 6
}

// Fan-Out, Fan-In

// These two patterns often work together. Fan-out means distributing work from one channel
// to multiple goroutines that process it in parallel.
// Fan-in (which is what Merge does) means collecting the results from those goroutines
// back into a single channel.
