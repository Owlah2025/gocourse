package main

import "fmt"

//type switch
/*func checkType(x interface{}){
switch x.(type) {
case int:
	fmt.Println("It's an int")
case string:
fmt.Println("it is a string ")
case float64:
fmt.Println("it is  a float")
default:
	fmt.Println("Unknown Type")

}
}*/
func main() {

//	x := 42.25
//using switch case to check type 
//checkType("hello")
//checkType(x)

x :=  5

switch x {
case 5:
	fmt.Println("x is: 5")
	fallthrough
case 4:
	fmt.Println("x is 4")
default:
fmt.Println("x is not found")

}


	
}
