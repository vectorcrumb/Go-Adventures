package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// We'll use a map for the state and a mutex for syncing access to the state
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	// These counters will track the amount of operations
	var readOps, writeOps uint64 = 0, 0

	// We'll then start 100 goroutines to repeatedly read the state, once per ms.
	// For each read, we pick a random key, lock the mutex, read the state and
	// unlock the mutex, before adding to the counter and sleeping the goroutine
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// We repeat the same idea but for write operations using 10 goroutines
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	// Finally, we report the counters and show the final state using a mutex lock
	fmt.Println("readOps:", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps:", atomic.LoadUint64(&writeOps))

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}