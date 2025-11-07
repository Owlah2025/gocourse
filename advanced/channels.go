package main

import "fmt"

func main() {
	//variable := make(chan type)

	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "World" //  you need to receive again
		//blocking because it is continuously trying to receive values,
		//it is ready to receive continuous flow of data

		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()
	reciever := <-greeting
	fmt.Println(reciever)
	// why not adding go ? because we are actualy in differnt goroutine.
	// main thread is a goroutine
	// but you can make another goroutine using go and communicate.

	reciever = <-greeting // same reciver? yeah it is fine
	fmt.Println(reciever)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	fmt.Println("End of program")

}
