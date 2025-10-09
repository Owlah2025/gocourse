package main

import "fmt"

func main() {
	
sequence := adder()
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())


}


func adder() func()int {
i := 0
fmt.Println("previous value of i:",i)

	return func ()int {
		i++
		fmt.Println("added 1 to i")
		return i
	}

}