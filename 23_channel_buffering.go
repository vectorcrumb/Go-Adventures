package main

import "fmt"

func main() {
	// By default, channels are unbuffered, meaning that sends are
	// only accepted if there's a corresponding receive ready to
	// receive the served value. Buffered channels accepted a
	// finite number of values without a corresponding receiver
	// to collect those values.

	// We make a channel with a buffer of size 2
	messages := make(chan string, 2)
	// We can add values into the channel without a concurrent receive
	messages <- "buffered"
	messages <- "channel"
	// Adding a third value to the channel raises a fatal error!
	//messages <- "woops!"
	// We can then receive those tho values in FIFO order
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
