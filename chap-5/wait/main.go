// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries += 1 {
		_, err := http.Head(url)

		if err == nil {
			return nil
		}

		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))

	}

	return fmt.Errorf("server %s faled to respond after %s", url, timeout)
}

func main() {
	url := "https://example.com"
	log.SetPrefix("wait: ")
	log.SetFlags(0)

	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}