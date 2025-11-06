package main

import "fmt"

func main() {
	
//NOTE: the adder() function is called once, but the anonymous function returned form it 
// //retains the value of i and called when we call the function variable (sequence). > closures

sequence := adder()
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())

fmt.Println("*******************")
anotherSequence := adder()
fmt.Println(anotherSequence())
fmt.Println(anotherSequence())
fmt.Println(anotherSequence())


subtracter := func () func(int) int  {
	countdown := 99
  return func(x int) int{
		countdown -= x
		return countdown
	}
}()

multiplier := func() func(int) int {
	factor := 2

	return func (x int) int  {
		return factor*x
	}


}()

fmt.Println(multiplier(5))
fmt.Println(multiplier(55))
fmt.Println("***************")
//using closure subtracter
fmt.Println(subtracter(3))
fmt.Println(subtracter(3))
fmt.Println(subtracter(3))
fmt.Println(subtracter(3))
fmt.Println(subtracter(3))

}


func adder() func()int {
i := 0 
fmt.Println("previous value of i:",i)

	return func ()int {
		i++
		fmt.Println("added 1 to i")
		return i
	}

}