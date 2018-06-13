package main

import (
	"math"
	"fmt"
)

/**
	There is no implicit type conversion in Go.
	You are supposed to use explicit type conversion
 */
func main() {
	triangle()
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
