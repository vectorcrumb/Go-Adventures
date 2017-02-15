package main

import (
	"time"
	"fmt"
)

func main() {
	// Suppose we wanted to limit the handling of incoming requests.
	// Let's create a channel to simulate the requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// We then create a channel to tick every 200 ms
	limiter := time.Tick(time.Millisecond * 200)
	// By receiving from the Tick channel, we block the handling
	// of requests by the specified time period. This is a basic
	// example of rate limiting
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// However, we could also want to receive bursts of requests
	// while preserving the overall rate limit. We accomplish
	// this by buffering the limiter channel.
	burstyLimiter := make(chan time.Time, 3)
	// We fill the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// Every 200 ms, we try to add a new value to
	// burstyLimiter up to its limit
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()
	// We then create 5 requests to be handled
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	// Finally, we handle the new requests. We'll
	// receive 3 requests in burst and then 2
	// more in 200 ms intervals
	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
