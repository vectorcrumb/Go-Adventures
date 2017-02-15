package main

import "fmt"

// This is a struct. It is a typed collection of fields,
// useful for grouping data together to form records.
type person struct {
	name string
	age int
}

func main() {
	// A struct can be created with inferred arguments
	fmt.Println(person{"Bob", 20})
	// The fields can be explicit when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})
	// Any omitted fields will be zero valued
	fmt.Println(person{name: "Fred"})
	// The & yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// The dot notation is used to access fields
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Pointers are automatically dereferenced when using dot notation on them
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
	s.name = "Bill"
	fmt.Println(sp.name)
}
