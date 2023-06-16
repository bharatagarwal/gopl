package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial( // returns net.TCPConn
		"tcp",
		"localhost:8000",
	)
	if err != nil {
		log.Fatal(err)
	}

	// copies from connection to stdout
	mustCopy(os.Stdout, conn)
}

func mustCopy(
	stdout *os.File, // implements io.Writer
	conn net.Conn, // implements io.Reader
) {
	// func Copy(dst Writer, src Reader) (written int64, err error)
	_, err := io.Copy(stdout, conn)
	if err != nil {
		log.Fatal( err)
	}
}
