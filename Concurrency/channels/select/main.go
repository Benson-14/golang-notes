package main

import (
	"fmt"
	"time"
)

// The select statement lets a goroutine wait on multiple channel operations at the same time
// Select evaluates all cases simultaneously and blocks until one of them can proceed

// An important difference from switch: if multiple cases are ready at the same time, select picks one at random
// A switch statement always evaluates cases top to bottom

// Timeout Pattern
// If a value arrives on ch within timeout, the first case executes
// If timeout expires, the second case executes
func WithTimeout(ch <-chan int, timeout time.Duration) (int, bool) {
	select {
	case msg := <-ch:
		return msg, true
	case <-time.After(timeout):
		return 0, false
	}
}

// TryReceive Pattern
// If a value arrives on ch, the first case executes
// If ch is empty, the second case executes
func TryReceive(ch <-chan int) (int, bool) {
	select {
	case msg := <-ch:
		return msg, true
	default:
		return 0, false
	}
}

func FanIn(ch1, ch2 <-chan int, done <-chan bool) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case msg := <-ch1:
				ch <- msg
			case msg := <-ch2:
				ch <- msg
			case <-done:
				return
			}
		}
	}()
	return ch
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 42
	val, ok := WithTimeout(ch1, 1*time.Second)
	fmt.Printf("WithTimeout (success): %d, %v\n", val, ok)

	ch2 := make(chan int)
	val, ok = WithTimeout(ch2, 100*time.Millisecond)
	fmt.Printf("WithTimeout (timeout): %d, %v\n", val, ok)

	ch3 := make(chan int, 1)
	ch3 <- 99
	val, ok = TryReceive(ch3)
	fmt.Printf("TryReceive (available): %d, %v\n", val, ok)

	ch4 := make(chan int)
	val, ok = TryReceive(ch4)
	fmt.Printf("TryReceive (empty): %d, %v\n", val, ok)
}
