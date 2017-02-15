package main

import (
	"fmt"
	"time"
)

func main() {
	// Timers represent a single event in the future. The constructor specifies
	// the wait time and it returns a channel that will be notified at that time.
	timer1 := time.NewTimer(time.Second * 2)
	// We then block the main routine until we receive from the channel timer.C
	<- timer1.C
	// If we wanted a delay, we could just use time.Sleep(). However, by using
	// a timer, we can stop the timer whenever we need to.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<- timer2.C
		fmt.Println("Time 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
