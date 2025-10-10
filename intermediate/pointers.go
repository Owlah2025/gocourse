package main

import (
	"fmt"
	"unsafe"
)

func main() {
	
	var a int = 10
	var ptr *int
	ptr = &a
	var ptr2 * int

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(ptr2)
	//NOTE: go doesn't support pointers arethmatic
	// go only allows referencing and dereferencing 

	
	modifyValue(ptr)
	fmt.Println(a)

	unSafePtr:=unsafe.Pointer(&ptr2) // converts the address of ptr2 to an unsafe.Pointer
	fmt.Println(unSafePtr)

}

func modifyValue(ptr *int){
	*ptr++
}

//NOTE: In Go, a function cannot be declared inside another function;
//only function literals (anonymous functions/closures) can be defined within a function body, 
// while named function declarations
