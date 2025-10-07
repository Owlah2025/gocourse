package main

import (
	"fmt"
	"slices"
)

func main() {
	
//var sliceName [] ElementType 

// var numbers [] int 
// var numbers1 =[]int {1,2,3}

// numbers2 := []int {9,8,7}
// slice :=make([]int,5)

a := [4]int{1,2,3,4} //this is an array not slice as we assigned the size 

mySlice:=a[:]

mySlice= append(mySlice,5,6,7,8)
fmt.Println("mySlice:",mySlice)

sliceCopy := make([]int,len(mySlice))
copy(sliceCopy,mySlice)
fmt.Println("sliceCopy:",sliceCopy)


fmt.Println("elments in sliceCopy are: ")
for i,v := range sliceCopy {
	fmt.Println("element at index",i,"is:",v)
}

if slices.Equal(sliceCopy,mySlice){
	fmt.Println("the two slices are equal")
}

twoDslice := make([][]int, 3)
for i:=0; i<3 ; i++ {
	twoDslice[i]=make([]int, i+1)
	for j:=0; j<i+1 ; j++{
		twoDslice[i][j] = i+j
	}


}
fmt.Println("twoDslice: ",twoDslice)

new2dslice := twoDslice[:][:]
fmt.Println("new2dSlice:",new2dslice)
fmt.Println("the capacity of new2dSlice: ",cap(new2dslice),"\nand the capacity of mySlice is:",cap(mySlice),"the length of mySlice",len(mySlice))

}