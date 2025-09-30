package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())  //source here is like the seed of 
	random := rand.New(source)                        //the random engine 

	target := random.Intn(100)+1        //genereatign a random number between 1 and 100
	
 fmt.Println("Welcome to the guessing game!! ")
 fmt.Println("I've choosen a number between 1 and 10.")
 fmt.Println("Can you guess the number ?")

var playerGuess int

 for {

	fmt.Println(" Enter your guess: ")
	fmt.Scanln(&playerGuess)

	if (playerGuess==target){
		fmt.Println("Bravooo!!! you guessed it right")	
		break

	} else if (playerGuess < target) {
		fmt.Println("number is too low! try again")
	}else {
		fmt.Println("the number is too high! try again")
	}

}
}



