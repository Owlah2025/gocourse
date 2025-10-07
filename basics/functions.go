package main

import "fmt"

func main() {
	
// 	//func <name> (parameters list) returnType {
// 	//code block
// 	//return value 

// 	//}

// //you can make this: 
// sum := add(2,7)
// fmt.Println(sum)
// //or print directly the funtion 
// fmt.Println(add(2,5))

// //Anonymous Function
// func(){
// 	fmt.Println("hello from anonymous function style 1")
// }() // this () is used to call the function 

// //another way is to assign this anonymous function to a variable like this:
// greeting := func(){
// 	fmt.Println("hello from anonymous function style 2")
// }
// greeting()
// // greeting()



//we are passing a function as an argument here 
result := applyOperation(4,3, add)
_ = result
fmt.Println("5+3 =",applyOperation(5,3,add))

//returning and using a function 
// muliplyBy3 := createMultiplier(3)
// multResult := muliplyBy3(3)
// fmt.Println("3*3 =",multResult)

muliplyBy3 := createMultiplier(3)
fmt.Println("3*4 =",muliplyBy3(4))

}






func add(a int, b int) int {
	return a + b
}


//Function that takes a function as an argument 
func applyOperation(x int , y int, operation func(int,int)int)int {
	return operation(x,y)
}

//Function that returns a function 
func createMultiplier(factor int) func (int) int {

	return func (x int) int {
		return x*factor
	}


}


//this is how we are proving that functions in go
//  is a first-class citizen or first-class object 