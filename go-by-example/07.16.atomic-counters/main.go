package main

import (
	"fmt"
	"sync"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i += 1 {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c += 1 {
				// atomic.AddUint64(&ops, 1)
				ops += 1
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
