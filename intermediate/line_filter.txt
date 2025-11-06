package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	//keyword to filter lines
	keyword := "important"

	//read and filter lines
	for scanner.Scan() {
		line := scanner.Text() //returns the string(s.token) wich is []byte
		if strings.Contains(line, keyword) {
			updateLine := strings.ReplaceAll(line, keyword, "neccessary")
			fmt.Printf("%d Filtered line: %v\n", lineNumber, line)
			fmt.Printf("%d Updated line: %v\n", lineNumber, updateLine)
			lineNumber++
		}
	}
	//remember that scanner may result in an error
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning filie:", err)
	}

}
