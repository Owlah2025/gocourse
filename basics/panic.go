package main

import "fmt"

func main() {
	
//panic(interface{})  interface is like any type of value 

//example of a valid input
process(10)

//example of an invalid input
process(-3)
	
}


func process(input int){

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0{
		fmt.Println("Before panic statemet")
		panic("input must be a non-negative number")
		// fmt.Println("After panic statement")   NOTE: this is unreachable code
		// defer fmt.Println("Deferred 3")        NOTE: even if it is deferred    
	}
	fmt.Println("Processing input:", input)

}

//as you see it is very powrfull using defer with panic 
//as it will not panic untill it executes the deferred statements 

//IMPORTANT: It panicked only after running the deferred functions or statemts