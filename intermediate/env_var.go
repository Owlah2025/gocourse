package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("USER") // NOTE: env vars are key value pairs
	home := os.Getenv("HOME")

	fmt.Println("user env var:", user)
	fmt.Println("Hme env var:", home)
	err := os.Setenv("FRUIT", "ORANGE")
	if err != nil {
		fmt.Println("Error setting new env var:", err)
	}
	fmt.Println("Env var FRUIT set successfully")
	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))

	// for _, envVar := range os.Environ() {

	// 	kvPair := strings.SplitN(envVar, "=", 2) //n > 0: at most n substrings;
	// 	//the last substring will be the unsplit remainder;
	// 	fmt.Println(kvPair[0]) //this will list all the keys (env vars)

	// }

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting env var:", err)
		return
	}
	fmt.Println("Unset env var done on key FRUIT")
	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))

	fmt.Println("------------------------------------------------------")
	//demonistrating SplitN() function
	str := "a=b=c=d=e"
	fmt.Println(strings.SplitN(str, "=", -1)) // this is the defualt; the same as Split()
	fmt.Println(strings.SplitN(str, "=", 0))  //zero substrings
	fmt.Println(strings.SplitN(str, "=", 1))  //one substring    // no seperation
	fmt.Println(strings.SplitN(str, "=", 2))  //two substrings
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))
	fmt.Println(strings.SplitN(str, "=", 5))

}
