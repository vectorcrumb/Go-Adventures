package main

import "fmt"
import "os"

func main() {
	// Defer is used to ensure that a function call is performed
	// later in a program's execution, usually for cleanup.
	// Immediately after opening the file, we defer the closing
	// of the file by calling closeFile. Finally, we continue to
	// work with the file, knowing that the deferred function
	// call will run at the end of the enclosing function, main
	f := createFile("files/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}