package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	// We can get the current time using time.Now()
	now := time.Now()
	p(now)
	// We can also construct a time struct with time.Date
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	// We can then extract various components of the time instance
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	// We can also ask for day of the week
	p(then.Weekday())
	// We can then compare between two time
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	// Using Sub, we can find the difference in between two dates, obtaining a Duration struct
	diff := now.Sub(then)
	p(diff)
	// We can also ask for components of a Duration
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	// With Add, we can advance a time by a given duration or reverse with a - sign
	p(then.Add(diff))
	p(then.Add(-diff))
}