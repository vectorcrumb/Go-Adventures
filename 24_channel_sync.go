package main

import (
	"fmt"
	"time"
)

// Channels may be used to sync execution across
// goroutines. Here, we use a blocking receive
// to wait for a goroutine to finish execution

// The argument syntax is *name* chan *type-val*.
// It creates a bool channel named 'done'
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	// Create a channel and give it to the subroutine
	done := make(chan bool, 1)
	go worker(done)
	// By using an empty call to receive from the channel,
	// we block main's execution until the worker is ready
	// and sends confirmation via the channel
	<-done
}
