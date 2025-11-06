package main

import (
	"fmt"
)

func main() {

	p1 := Person{
		firstName: "Ahmed",
		lastName:  "Hazem",
		age:       25,
	}

	p2 := Person{
		firstName: "Ali",
	}
	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println("the full name is:", p1.fullName())

	//Anonymous Struct
	user := struct {
		userName string
		email    string
	}{
		userName: "AboAli",
		email:    "aboali@gmail.com",
	}

	fmt.Println(user)

	fmt.Println("Before increment p1.age:", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment p1.age:", p1.age)
	p1.age.increment()
	fmt.Println("After another increment using type Age:", p1.age)
}

type Age int //IMPORTANT NOTE : you have to make a type for age to allow you
// to make a method on it
func (a *Age) increment() {
	*a++
}

type Person struct {
	firstName string
	lastName  string
	age       Age
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

//METHODS
//func (method_reciver) method_name(){
//
//}
