// Worker Pools
// - Using goroutines & channels

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker num:", id, "started job:", j)
		time.Sleep(time.Second)
		fmt.Println("Worker num:", id, "finished job:", j)
		results <- j * 2
	}
}

func main() {
	// Var for 5 jobs
	const numJobs = 10

	// Set up 2 channels, one for sending jobs, one for collecting the results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start up 3 workers (no jobs sent so its blocking)
	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs then close it to say we've sent all jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Now collect the results (will print all results out)
	for r := 1; r <= numJobs; r++ {
		<-results
	}
}