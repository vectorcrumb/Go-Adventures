package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// The most basic way of reading files is to transfer the whole content of a file into memory
	dat, err := ioutil.ReadFile("files/data")
	check(err)
	fmt.Println(string(dat))

	// However, we'll usually want more control over the file. For this purpose, we'll use os.Open to create an os.File
	f, err := os.Open("files/data")
	check(err)

	// Read bytes from the beginning. Allow up to 5, but also take note of how many were read
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// Seek to jump to a known location and read from there
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// With the io package we get more robust tools for operations like the previous one
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err :=  io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no rewind, but we can Seek back to (0, 0)
	_, err = f.Seek(0, 0)
	check(err)

	// bufio implements a buffered reader, efficient for small reads and with more functions
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file when done. This would usually be done with a defer call after opening the file
	f.Close()
}