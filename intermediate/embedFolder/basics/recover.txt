package main

import "fmt"

func main() {
	
	process()
	fmt.Println("Returned from porcess!!")
}

func process(){

	defer func(){
		// r:=recover()
		// if  r!=nil{
		// }                              this is the same as the following:
		if r:=recover(); r!=nil{
					fmt.Println("Recovered:",r)
		}
	}()


	fmt.Println("Start Process!!!!")
	panic("Somthing went WRONG!!")
	fmt.Println("after panic statement 'unreachable'")

}