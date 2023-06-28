package main

import (
	"fmt"
	"time"
)

func main() {
	// start := time.Now().UnixMilli()
	timer1 := time.NewTimer(2 * time.Second)

	// receive firing event in 2 seconds
	<-timer1.C // blocking
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(1 * time.Second)

	go func() {
		<-timer2.C // blocking event in this goroutine
		fmt.Println("Timer 2 fired")
	}()

	// cancel before firing
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
