package main

import "crypto/sha1"
import "fmt"

// Go implements several hash functions in various crypto packages

func main() {
	s := "sha1 this string"
	// To generate a hash, we follow 3 steps.
	// First, we create a new hash generator
	h := sha1.New()
	// Then we call Write on the generator with a slice of bytes.
	// We have to convert our string to a slice of bytes
	h.Write([]byte(s))
	// Finally, we call Sum on the generator. The argument to
	// sum simply appends mores bytes to the slice stored, but
	// it generally isn't used
	bs := h.Sum(nil)
	// SHA1 values are printed in hex, so we format with %x
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
