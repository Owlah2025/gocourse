package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// var stringFlag string
	// flag.StringVar(&stringFlag, "user", "Ahmed", "the name of the user")

	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subcommand1.Bool("processing", false, "Command processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subcommand2.String("language", "Go", "Enter your language")
	// flag.Parse()
	// fmt.Println("user name:", stringFlag)

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("processing:", *firstFlag) //NOTE: remeber to use * as bool or string returns adress of a variable
		fmt.Println("bytes:", *secondFlag)

	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language:", *flagsc2)

	default:
		fmt.Println("no subcommand entered!")
		os.Exit(1)

	}

	//go run command_line_sub_commands.go firstSub -Processing=true -bytes=256
	// subCommand1:
	// processing: true
	// bytes: 256

	//	//ahmed-hazem@AHazem:~/Documents/gocourse/intermediate$ go run command_line_sub_commands.go firstSub --help
	//
	// Usage of firstSub:
	//
	//	-bytes int
	//	      Byte length of result (default 1024)
	//	-processing
	//	      Command processing status
}
