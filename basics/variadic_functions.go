package main

import "fmt"

func main() {
	// ... these 3 dots is called Ellipsis 

	// func functionName(param1 type1, param2 ...type2) returnType {
	// 	function body 
	// } 

	// fmt.Println("sum of 1,2,3 : ", sum(1,2,3))
	statment , total := sum("the some of numbers is:",3,4,7,5 ,6 ,25)
	fmt.Println(statment,total)

	numSlice := []int {3,3,52,252,5}

	fmt.Println(sum("the sum of the slice is ",numSlice...)) //NOTE you can use slices as variadic parameters
																													// it is destructuring 
}


func sum(returnString string, nums ...int) (string, int){
total := 0
for _,v := range nums {
	total +=v
}
return returnString, total

}


//NOTE:if a function accepts variadic parameter it is called variadic function
// IMPORTANT > all variadic parameters must be the last parameters in the function signature 