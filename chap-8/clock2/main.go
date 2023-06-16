package main

import (
	"io"
	"log"
	"net"
	"time"
)

// handles one complete client connection
func handleConn(conn net.Conn) {
	// close connection before returning
	defer conn.Close()

	for {
		_, err := io.WriteString(
			conn,
			time.Now().Format("15:04:05\n"),
		)

		// end loop when write has failed
		// most likely because client has disconnected.
		if err != nil {
			break
		}
		time.Sleep(time.Second)
	}
}

func main() {
	// listens for incoming connections.
	listener, err := net.Listen( // returns a net.Listener
		"tcp",
		"localhost:8000",
	)

	if err != nil {
		log.Fatal(err)
	}

	// a server is expected to keep running until a user closes it explicitly.
	for {
		// blocks until connection requests is made
		conn, err := listener.Accept() // returns a net.Conn object

		// Print if there is an error with connection
		// Keep loop going
		if err != nil {
			log.Print(err)
			continue
		}

		// wait for connection to be handled before moving to next connection.
		// Adding go to this will let next connection be accepted.
		go handleConn(conn)
	}
}