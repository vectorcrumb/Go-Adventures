package main

import "fmt"

func main() {
	// We'll use the jobs channel to communicate
	// jobs from the main routine to a goroutine
	jobs := make(chan int, 5)
	done := make(chan bool)

	// The anonymous worker goroutine will continuously receive values from
	// the jobs channel. The double return value syntax returns a message
	// followed by a bool indicating if the channel isn't empty.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job:", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	// We then send 3 jobs over to the worker
	for j := 1; j <=3; j++ {
		jobs <- j
		fmt.Println("sent job:", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	// We wait using the sync approach seen before. If we didn't sync like this,
	// the program would end prematurely, killing the worker before it's finished
	<- done
}
