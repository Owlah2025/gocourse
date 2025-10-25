package main

import "fmt"

func main() {

	process(5)

}

func process(i int){
	//IMPORTANT: last deferred first executed.!!! LIFO > to ensure the right resource cleanup order
	
	defer fmt.Println("Defrred i value1:",i)
	defer fmt.Println("First deferred statment executed")
	defer fmt.Println("Second deferred statment executed")
	defer fmt.Println("Third deferred statment executed")
	i++
	fmt.Println("Normal Excution Statment")
	fmt.Println("Value of i:",i)
	defer fmt.Println("Defrred i value2:",i)
	//defer is a verb that means أجِّل
}
//[]