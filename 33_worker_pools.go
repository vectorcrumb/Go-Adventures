package main

import (
	"fmt"
	"time"
)

// Create a worker function with two unidirectional channels
func workerp(id int, jobs <-chan int, results chan<- int) {
	// Go through all the jobs and pipe back the results
	for j:= range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// Create a channel for jobs and another for results
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// Spawn 3 workers who'll block until they receive jobs
	for w := 1; w <= 3; w++ {
		go workerp(w, jobs, results)
	}
	// Send jobs over to the works. Workers will not receive
	// 5 jobs each, but instead process jobs as they finish
	// and there still are jobs in the channel
	for j := 1; j <=5; j++ {
		jobs <- j
	}
	close(jobs)
	// Collect the results
	for a := 1; a <= 5; a++ {
		<- results
	}
}