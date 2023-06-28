package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println("Worker", id, "starting")
	time.Sleep(time.Second)
	fmt.Println("Worker", id, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i += 1 {
		wg.Add(1)

		i := i
		// fixing value in place for closure

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
