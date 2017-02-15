package main

import "fmt"

// Channel directions may be explicit.
// This increases the type-safety of the code

// pings is a send-only channel. Trying to receive
// from send would cause a compile-time error
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong takes two channels: pings for receive-only
// and pongs for send only. Both are string channels
func pong(pings <-chan string, pongs chan<- string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	// Two channels are created
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	// A message is passed onto pings with ping()
	ping(pings, "passed message")
	// The message from pings is passed onto pongs with pong()
	pong(pings, pongs)
	// The message inside pongs is extracted and printed out
	fmt.Println(<-pongs)
}
