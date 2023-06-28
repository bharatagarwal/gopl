package main

import (
	"fmt"
	"time"
)

// three workers are present
// trying to catch work when
// it is presented to jobs channel
func worker(id int, jobs <-chan int, results chan<- int)  {
    for value := range jobs {
        fmt.Println("worker", id, "started job", value)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", value)
        results <- value
    }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // starting three workers
    for w:= 1; w <= 3; w += 1 {
        go worker(w, jobs, results)
    }

    // workers are blocked, waiting for work
    // send values from 1 to 5 to jobs channel
    for j := 1; j <= numJobs; j += 1 {
        jobs <- j
    }
    
    // indicate to jobs channel that it can stop looking
    close(jobs)

    // collect results, draining results channel
    // iterating over numJobs ensured that all
    // goroutines are closed
    for a := 1; a <= numJobs; a += 1 {
        fmt.Println("result for job", <-results)
    }
}
