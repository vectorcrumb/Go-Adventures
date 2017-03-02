package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// We can create a simple command which takes no input
	// and prints to stdout. The exec.Command function
	// creates a object to represent this external process
	dateCmd := exec.Command("date")
	// With .Output, we execute the command, block until it finishes
	// and collect the output, checking for errors in the process
	dateOut, err := dateCmd.Output()
	if (err != nil) {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Now, we'll see a more complicated example where pipe data to
	// an external process on stdin and collect output from stdout
	grepCmd := exec.Command("grep", "hello")
	// We explicitly grab the pipes
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	// Then, we start the process and write some input into it
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	// We close the input pipe, read the output and wait for the process to exit
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	// Finally, we show the output. We have omitted the error checking above.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// When spawning commands, we must pass all arguments as an argument array
	// instead of a single string. To avoid this, we can use bash's -c flag
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}