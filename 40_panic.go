package main

import "os"

// A panic is used for more severe or unexpected errors. We'll use
// them to fail fast on errors that shouldn't occur on normal
// operation or those we aren't prepared to handle gracefully

func main() {
	// This line raises a panic with the error string
	panic("a problem")
	// A common use of panic is to abort if we receive a
	// error code we don't want to/know how to handle
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}