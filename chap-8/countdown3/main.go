package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Lift off!")
}

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)

	for i := 10; i > 0; i -= 1 {
		fmt.Println(i)

		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}

	}

	launch()
}