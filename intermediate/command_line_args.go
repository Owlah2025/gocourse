package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// //NOTE:  go clean -cache command make a force rebuild.
	fmt.Println("command0:", os.Args[0])
	fmt.Println("command1:", os.Args[1])

	for i, arg := range os.Args {
		fmt.Println("Argument number:", i, ":", arg)
	}

	// define flags
	var name string
	var age int
	var male bool
	flag.StringVar(&name, "name", "Ahmed", "the name of user")
	flag.IntVar(&age, "age", 25, "the age of the user")
	flag.BoolVar(&male, "gender", true, "the gender of the user")
	flag.Parse() //must be called afeter all flags are defined and before flags are accessed by the program

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("male:", male)
	// go run command_line_args.go -name Ahmed Hazem -age 23
	// go run command_line_args.go  Ahmed Hazem  23  // defualt value will take place without flags used

}
