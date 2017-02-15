package main

import (
	"fmt"
	"time"
	"sync/atomic"
)

func main() {
	// We'll use an unsigned int for a counter
	var ops uint64 = 0
	// We then start 50 goroutines to increment
	// the counter once per millisecond
	for i := 0; i < 50; i++ {
		go func() {
			for  {
				// AddUint64(pointer to uint64, delta)
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// We then wait a second to let the goroutines count
	time.Sleep(time.Second)
	// Using LoadUint64(pointer to uint64) we can safely
	// extract the value of ops counter, even while it's
	// being mutated by other concurrent goroutines
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}