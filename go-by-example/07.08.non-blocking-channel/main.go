package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking
	select {
	// messages is empty, msg is waiting -- blocking
	case msg := <-messages:
		fmt.Println("received message", msg)
	default: // saves the day
		fmt.Println("no message received")
	}

	msg := "hi"

	// non-blocking
	select {
	// no receiver for this channel -- blocking
	case messages <- msg:
		fmt.Println("received message", msg)
	default: // saves the day
		fmt.Println("no message received")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}