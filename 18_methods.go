package main

import "fmt"

type rect struct {
	width, height int
}

// This function has what is called a receiver type. In this case, it is of type *rect.
// A pointer receiver is used to avoid copying on method calls and can modify the original argument
func (r *rect) area() int {
	return r.width * r.height
}

// This method instead has a receiver type of rect.
// A value receiver receives a copy of the struct and cannot mutate the receiving struct
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Conversions between values and pointers is automatically handled for method calls
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
