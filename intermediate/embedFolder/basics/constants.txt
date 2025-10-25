package main

import (
	"fmt"
	"math"
)

const PI = 3.14
const GRAVITY = 9.81

func main() {
var smallFloat float64 = 1.0e-323
fmt.Println(smallFloat)
smallFloat/=math.MaxFloat64
fmt.Println(smallFloat)

}

