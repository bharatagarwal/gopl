package main

import "fmt"

// In the pipeline below, when the counter goroutine
// finishes its loop after 100 elements, it closes the
// naturals channel, causing the squarer to finish its loop
// and close the squares channel. (In a more complex
// program, it might make sense for the counter and squarer
// functions to defer the calls to close at the outset.)
// Finally, the main goroutine finishes its loop and the
// program exits.

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x += 1 {
			naturals <- x
		}

		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}