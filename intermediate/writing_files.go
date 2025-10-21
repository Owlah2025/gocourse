package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	//closing file
	defer file.Close() // using defer will close the file at the end of the main
	_, err = file.Write([]byte("Hello World!\n"))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data has been written to file successfully.")

	file, err = os.Create("writeString.txt") // we don't use the gopher assignment as these variables declared before
	if err != nil {
		fmt.Println("Error creating writeString file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString("Hello Go!\n\n\n\n\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Writing to writeString file complete.")
}
