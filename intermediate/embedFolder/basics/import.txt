package main

import (
	"fmt"
	"math"
	foo "net/http"
)


func main(){

	fmt.Println("Hello, Go standard library")
	resp,err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status: ", resp.Status)
	fmt.Println(math.Abs(-924));
	
	
}