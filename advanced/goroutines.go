package main

import (
	"fmt"
	"time"
)

//Goroutines are just functions that leave the main thread and run in the background
//and come back to join the main thread once the functions are finished/ready to retrun any value
//Goroutines do not stop the program flow and are non blocking

func main() {

	var err error

	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After sayHello function")

	//err = go doWork() // this is not accepted
	go func() {
		err = doWork()
	}()

	go printLetters()
	go printNumbers()

	time.Sleep(2 * time.Second) // prevent Goroutine leaks

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("work completed successfully")
	}
}

// ///////////////////////////////////////////////////////////////////////////////
func sayHello() {
	time.Sleep(time.Second * 1)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := range 5 {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("letter:", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occured in doWork.")
}

// non-deterministic nature of goroutine scheduling in Go
//Excessive creation of go routines can lead to resource exhaustion
