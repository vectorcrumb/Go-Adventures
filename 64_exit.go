package main

import "fmt"
import "os"

func main() {
	// defers are NOT run when exiting!
	defer fmt.Println("!")
	// exit with a status
	os.Exit(3)
}
