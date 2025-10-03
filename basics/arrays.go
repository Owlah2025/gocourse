package main

import (
	"fmt"
)

func main() {

	myArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArr)

	myArr[1] = 22
	fmt.Println(myArr)

	copiedArr:= myArr

	for i,_ := range myArr{  //you can simplify: just remove the secod part _ and use i only
		fmt.Println(i)
	}
	fmt.Println("***********************")

	for i:=0; i<len(copiedArr);i++ {
		fmt.Println(copiedArr[i])
	}
	fmt.Println("***********************")

	for i:=range len(copiedArr) {												//the same as above
		fmt.Println(copiedArr[i])
	}

	fmt.Println("***********************")

	a,b := aFunction()
	fmt.Println("first number is:",a,"\nthe second number is:",b)
}

//function to return 2 values 
func aFunction() (int,int){
	return 2,3
}