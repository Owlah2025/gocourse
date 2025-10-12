package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64 // you have to implement these methods ...it is a contract!!
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
	fmt.Println(c.diameter())
	myPrinter(42, 24.2, "Hello", 0b10101, "a", []any{4, 2325, 'a'}, true) //any type
	value(2235)
	printType(true)
}

// **************************************************************************
//
//	func myPrinter(i ...interface{}) {
//		for _, v := range i {
//			fmt.Println(v)
//		}
//	}
func myPrinter(i ...any) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func value(anyType any) {
	fmt.Println(anyType)
	// fmt.Println(anyType.(type)) IMPORTANT NOTE: you can not use .(type) outside type switch
}

func printType(i any) {
	switch i.(type) {

	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	default:
		fmt.Println("Unknown Type")

	}
}
