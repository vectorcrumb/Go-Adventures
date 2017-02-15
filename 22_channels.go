package main

import "fmt"

func main() {
	// Channels are pipes which connect concurrent goroutines.
	// You can send values in and receive them in another goroutine

	// New channels are created with make(chan val-type)
	messages := make(chan string)
	// Send a value into the channel using channel <- value
	go func() {
		messages <- "ping"
	}()
	// Receive a value from a channel using value <- channel
	msg := <- messages
	fmt.Println(msg)
	// By default, sends and receives block until the sender and
	// receiver are both ready. This allows us to receive the
	// "ping" message without additional synchronization
}
