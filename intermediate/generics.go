package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	poppedElement := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return poppedElement, true
}

func (s Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty man!!")
		return
	}
	fmt.Println("Stack elemets:")
	for _, element := range s.elements {
		fmt.Println(element)
	}
}

func main() {

	// /*
	// 	x, y := 1, 3
	// 	x1, y1 := "Ahmed", "Hazem"

	// 	x, y = swap(x, y)
	// 	x1, y1 = swap(x1, y1)

	// 	fmt.Println(x, y)
	// 	fmt.Println(x1, y1)

	// */

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Print("popped: ")
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Print("popped: ")
	fmt.Println(intStack.pop())
	fmt.Println(intStack.isEmpty())

	fmt.Print("popped: ")
	fmt.Println(intStack.pop())
	fmt.Print("popped: ")
	fmt.Println(intStack.pop())

	intStack.printAll()
	fmt.Println(intStack.isEmpty())

}
