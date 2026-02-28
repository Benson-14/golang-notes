package main

import "fmt"

func generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range n {
			out <- i
		}
		close(out)
	}()

	return out
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func main() {

	ch := make(chan int) // Create a channel that carries int values

	go func() {
		ch <- 42 // Send the value 42 into the channel
	}()

	value := <-ch // Receive a value from the channel
	fmt.Println(value)

	// Notice that we launch the send in a goroutine. This is necessary because an unbuffered channel
	// (created without a size) requires both a sender and receiver to be ready
	// at the same time. If we tried to send and receive in the same goroutine
	// without a buffer, the program would deadlock

	// value will be printed only after it receives something from the channel

	// CLOSING CHANNELS
	// Only the sender should close the channel
	// Sending on a closed channel causes a panic

	go func() {
		ch <- 1
		ch <- 2
		ch <- 100
		close(ch) // Signal that we're done sending
	}()

	for v := range ch {
		fmt.Println(v) // Prints 1, then 2, then the loop exits
	}

	// Generator Pattern
	for num := range generate(5) {
		fmt.Printf("%d ", num)
	}

	// The function returns a receive-only channel (<-chan int). The caller can only read from it.
	for sq := range Square(generate(5)) {
		fmt.Printf("%d ", sq)
	}

}
