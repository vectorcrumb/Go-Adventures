package main

import "fmt"

func main() {
	// We open a buffered channel of size two and then close it. Closing
	// a non-empty, buffered channel doesn't remove the values inside it
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// We can use for and range to iterate over values received in a channel.
	// Because the channel was previously closed, the iteration stops after
	// receiving both stored values. If the channel hadn't been close, we
	// would've received a fatal error
	for elem := range queue {
		fmt.Println(elem)
	}
}
