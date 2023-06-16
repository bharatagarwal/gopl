package main

import "fmt"

// The first goroutine, counter, generates the integers 0,
// 1, 2, ..., and sends them over a channel to the second
// goroutine, squarer, which receives each value, squares
// it, and sends the result over another channel to the
// third goroutine, printer, which receives the squared
// values and prints them. For clarity of this example, we
// have intentionally chosen very simple functions, though
// of course they are too computationally trivial to warrant
// their own goroutines in a realistic program.

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x += 1 {
			naturals <- x // send operation
		}
	}()

	go func() {
		for {
			x := <-naturals  // receive operation
			squares <- x * x // send operation
		}
	}()

	for {
		fmt.Println(<-squares) // receive operation
	}
}