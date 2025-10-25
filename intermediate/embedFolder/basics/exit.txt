package main

import (
	"fmt"
	"os"
)

func main() {
	
	fmt.Println("Starting the main function")
	defer fmt.Println("Deferred statement")

	//Exit with status code of 1
	os.Exit(1) // NOTE It's and immediate exit without running any cleanup operations

	//This will never be executed
	fmt.Println("End of main function")



}