package main

import "fmt"

//IMPORTANT: Why init Works Multiple Times?????
//The Go compiler treats init as a special function for package initialization.

//You can declare func init() { ... } as many times as you want in a package or file, and all will be executed in order before main.


func init(){
	fmt.Println("Initializing package1...")
}
func init(){
	fmt.Println("Initializing package2...")
}
func start(){
	fmt.Println("Initializing package3...")   //NOTE: this function failed to run 
}


func main() {
	fmt.Println("Inside the main function")
}