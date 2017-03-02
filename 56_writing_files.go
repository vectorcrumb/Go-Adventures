package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check2(e error) {
	if e != nil {
		panic(e)
	}
}

func main () {
	// We can simply dump a string into a file
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("files/data1", d1, 0644)
	check2(err)
	// For more specificity, we can open a file for writing
	f, err := os.Create("files/data2")
	check2(err)
	// We immediately defer the closing of the file
	defer f.Close()

	// Writing byte slices is forward
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check2(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// We can also use WriteString
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// With Sync, we flush writes to a stable storage from program memory
	f.Sync()

	// bufio provides buffered writers in addition to the readers
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check2(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// With Flush on a buffered writer we ensure all buffered
	// operations are applied to the underlying writer
	w.Flush()
}