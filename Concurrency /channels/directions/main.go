package main

import "fmt"

// Channel Direction Types
// chan T        // Bidirectional - can send and receive
// chan<- T      // Send-only - can only send
// <-chan T      // Receive-only - can only receive

// When a function parameter is chan<- T, the function can only send values into the channel.
// When a function parameter is <-chan T, the function can only receive values. It cannot send or close the channel

func Produce(ch chan<- int, n int) {
	// Send 1 to n, close when done
	for i := range n {
		ch <- i + 1
	}
	close(ch)
}

func Consume(ch <-chan int, result chan<- int) {
	// Sum all values, send sum to result
	sum := 0
	for v := range ch {
		sum += v
	}
	result <- sum
}

func Transform(in <-chan int, out chan<- int) {
	// Multiply by 10, close when done
	for v := range in {
		out <- v * 10
	}
	close(out)

}

func main() {
	// Test Produce -> Consume
	ch := make(chan int)
	result := make(chan int)

	go Produce(ch, 5)
	go Consume(ch, result)

	fmt.Println("Sum of 1-5:", <-result) // 15

	// Produce -> Transform -> Consume
	ch1 := make(chan int)
	ch2 := make(chan int)
	result2 := make(chan int)

	go Produce(ch1, 5)
	go Transform(ch1, ch2)
	go Consume(ch2, result2)

	fmt.Println("Sum of (1-5)*10:", <-result2) // 150
}
