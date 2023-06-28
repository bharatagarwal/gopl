package main

import (
	"fmt"
	"sync"
)

// Container has a map and a mutex
// Mutex: state int32, sema  uint32
// The zero value for a Mutex is an unlocked mutex.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name] += 1
}

func main() {
	c := Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i += 1 {
			c.inc(name)
		}

		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}