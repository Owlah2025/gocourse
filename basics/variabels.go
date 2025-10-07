package main

import "fmt"

var middleName string = "Hazem"
var emptyString string 

func main(){

	var name = "Ahmed"
	var lastName = "Elmelegy"
	fmt.Println(name)
	fmt.Println(lastName)
	
	var age int = 24
	
	num :=5235

	fmt.Println(age)
	fmt.Println(num)
	
personal()

fmt.Println(middleName,emptyString)

}

func personal(){
	firstName:="Ahmed"
	fmt.Println(firstName)
}