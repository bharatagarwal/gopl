package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i += 1 {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// limiter becomes blocking
		// until it receives from Tick
		// every 200ms
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
