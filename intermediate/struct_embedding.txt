package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	//employeeInfo person //Embededed struct Named field
	person //Anonymous field   //NOTE: field promotion works only if it is anonymous
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Println("the name is:", p.name, "the age is:", p.age)
}
func (e employee) introduce() {
	fmt.Println("Hi I am", e.name, "employee ID:", e.empId, "and I earn", e.salary)
}

func main() {
	emp := employee{
		person: person{name: "Ahmed", age: 25},
		empId:  "E001",
		salary: 25.35,
	}

	//field promotion
	fmt.Println("Name:", emp.name) //instead of accessing it emp.person.name
	fmt.Println("Age:", emp.age)
	fmt.Println("Employee id:", emp.empId)
	fmt.Println("Salary:", emp.salary)
	p := person{name: "Ali", age: 22}
	emp.introduce()
	p.introduce()

}
