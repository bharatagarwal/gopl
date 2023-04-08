package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
Copy returns the byte count, along with any error that occurred.
As each result arrives, fetch sends a summary line on the channel ch. The second
range loop in main receives and prints those lines.
*/

func main() {
	start := time.Now()
	// The main function creates a channel of strings using make.
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	// The second range loop in main receives and prints those lines.
	for range os.Args[1:] {
		/*
			When one goroutine attempts a send or receive on a channel,
			it blocks until another goroutine attempts the corresponding
			receive or send operation, at which point the value is
			transferred and both goroutines proceed.

			In this example, each fetch sends a value (ch <- expression)
			on the channel ch, and main receives all of them (<-ch).

			Having main do all the printing ensures that output from each
			goroutine is processed as a unit, with no danger of interleaving
			if two goroutines finish at the same time.
		*/

		fmt.Println(<-ch) // receive from channel
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel
		return
	}

	// Copy returns the byte count, along with any error that occurred.
	nbytes, readingError := io.Copy(io.Discard, resp.Body)

	closingError := resp.Body.Close() // don't leak resources

	if closingError != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		if readingError == nil {
			return
		}
	}

	if readingError != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, readingError)
		return
	}

	secs := time.Since(start).Seconds()

	// As each result arrives, fetch sends a summary line on the channel ch.
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
