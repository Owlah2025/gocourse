package main

import (
	"errors"
	"fmt"
)

func main() {

q,r := divide(10,3)
fmt.Println("Quotient:",q,"\nRemainder:",r)

result,err := compare(0,0)
if err !=nil {
	fmt.Println("Error:",err)
}else{
	fmt.Println(result)
}

}

//Multiple return value is important in error observation 
// func divide(a int , b int) (int, int){

// 	quotient := a/b
// 	remainder := a%b 

// 	return quotient, remainder
// }

// some adjustments 
func divide(a int , b int) (quotient int,remainder int){

	quotient = a/b      //don't use gopher like before 
	remainder = a%b 		//because it is already declared
	return							//you can ommit the statments after return keyword
}											//go compiler is smart enough 

func compare(a int ,b int) (string,error){

	if a > b {
		return "a is greater than b",nil
	}else if b > a {
		return "b is greater than a", nil 
	}else {
		return "",errors.New("Unable to compare which is greater")
	}
}




//char is called rune in go 
// you can iterate over strings using range 
// message := "hello world"
// for i , v :=range message { 
// fmt.Printf("Index: %d, Rune: %c\n",i,v) 
//
//}