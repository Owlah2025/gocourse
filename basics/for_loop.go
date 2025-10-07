package basics

import "fmt"

func main() {
	
// for i:=1 ; i<=5; i++{
// fmt.Println(i)

// }
// var value = 9
// fmt.Printf("value: %d\n", value)

var rows = 5 

for i:=1; i<=rows; i++ {
	for j:=1; j<=rows-1; j++ {
		fmt.Print(" ")
	}
	for k:=1; k<=i;k++ {
		fmt.Print("*")
	}
	fmt.Println()

	for i:= range 12{
		i++
		fmt.Println("hello")
	}
	
}


}