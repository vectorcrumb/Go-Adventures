package main

import "fmt"
import "time"

func main() {
	// Unix or UnixNano on a Now object gives the seconds or nanoseconds since the epoch
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	// There is no UnixMillis, so to get the ms, divide the ns by 1000000
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	// You can convert ms or ns since the epoch as ints to time
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
