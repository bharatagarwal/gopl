package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial(
		"tcp",
		"localhost:8000",
	)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		// ignoring error because closing the closing the connection
		// will lead to a read from closed connection error
		io.Copy(
			os.Stdout,
			conn,
		)
		fmt.Println("done")
		done <- struct{}{} // send to done channel
	}()

	mustCopy(conn, os.Stdin) // can be closed by client entering Ctrl+D
	conn.Close()
	<-done // receive from done channel
}

func mustCopy(
	conn net.Conn,
	stdin *os.File,
) {
	_, err := io.Copy(conn, stdin)
	if err != nil {
		log.Fatal(err)
	}
}