// netcat2
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// retrieving output from connection
	go mustCopy(os.Stdout, conn)
	
	// sending string to be echoed to connection
	mustCopy(conn, os.Stdin)
}


// use interfaces because operands are switched 
// in line 18 and 19
func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}