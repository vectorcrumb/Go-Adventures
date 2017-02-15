package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key 	int
	resp 	chan int
}

type writeOp struct {
	key 	int
	val 	int
	resp 	chan bool
}

func main() {
	// Counters to keep track of operations
	var readOps, writeOps uint64 = 0, 0

	// We've created two channels to transmit read/write operations.
	// They're pointer channels so the goroutine in charge of state
	// can simply write to the internal channel with a response.
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// This goroutine is owner of the state, which is local to it.
	// It loops a select which is in charge of responding to either
	// read or write requests on both channels
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// We then start 100 goroutines to issue reads to the state
	// goroutine. Each iteration requires constructing a readOp,
	// sending it over the reads channel and obtaining the result
	// via rhe read.resp channel.
	for r := 0; r < 100; r++ {
		go func() {
			for  {
				read := &readOp{
					key: rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// Same as before, but 10 goroutines for writing
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	// Finally, we report the operation counts
	fmt.Println("readOps:", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps:", atomic.LoadUint64(&writeOps))
}
