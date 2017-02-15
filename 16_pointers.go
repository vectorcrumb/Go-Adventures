package main

import "fmt"

// zeroval takes an int argument and sets it to zero.
// However, values are copied from the argument to a
// local copy, so the original value is not affected
func zeroval(ival int) {
	ival = 0
}

// zeroptr takes a reference to an int as an argument,
// dereferences the pointer reference and sets the
// value at the referenced address to zero
func zeroptr(iptr *int) {
	*iptr = 0
}

/*
The syntax then is as follows:
*type: As an argument, it indicates the function receives a pointer.
*variable: As a statement, it takes a pointer variable and dereferences it.
&variable: As a statement, it returns a pointer to the variable's memory address
 */
func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
