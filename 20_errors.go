package main

import (
	"errors"
	"fmt"
)

// By convention, errors are tha last return value and use the error interface
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New creates a basic error
		return -1, errors.New("Can't work with 42")
	}
	// A nil value for error indicates that there was no error
	return arg + 3, nil
}

// We can use custom types as errors by implementing Error() on them
type argError struct {
	arg int
	prob string
}

// Implemented Error() with a pointer receiver
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// This function uses the custom error, returning a &argError in case of error
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// Test f1 with {7, 42}
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	// Test f2 with {7, 42}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	// Test f2 with 42. This causes an error. The error is picked up in e
	// and then cast as an instance of the custom error type via type assertion.
	// The pointer in the type assertion is due to the Error() implementation
	// using a pointer receiver. That is why the error returned in f2 is a
	// pointer to the new error's mem address.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}