package main

import (
	"syscall"
	"os"
	"os/exec"
)

func main() {
	// First, we find the absolute path for the ls binary
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	// Exec requires arguments in slice form, not as one string. The first value must be the program name
	args := []string{"ls", "-a", "-l", "-h"}
	// Finally, exec requires a set of environment values
	env := os.Environ()
	// With these 3 variables, we can make a syscall.Exec call. If the call is successful, the execution
	// of this program ends and it will be replaced by the process specified in the previous variables.
	// In conclusion, we effectively replace the current Go process with a different process.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
