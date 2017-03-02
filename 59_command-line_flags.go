package main

import "flag"
import "fmt"

func main() {
	// Basic flag declarations are available for strings, integers and
	// booleans. Here, we've declared a string flag with a default value
	// of foo. The flag.String function returns a pointer to a string
	wordPtr := flag.String("word", "foo", "a string")
	// We repeat the same for boolean and int types
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	// We can also use an existing variable to store values.
	// We must give a pointer to the variable
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// With all flags declared, call Parse to receive the command-line inputs
	flag.Parse()

	// Remember to dereference the pointers when getting the corresponding values
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	// flag.Args() contains additional trailing arguments, as in os.Args.
	// All flags MUST appear before trailing positional arguments
	fmt.Println("tail:", flag.Args())

	// Calling the -h flag returns a help prompt with the
	// descriptions and default values specified before

}
