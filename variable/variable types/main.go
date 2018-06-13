package main

import (
	"math/cmplx"
	"math"
	"fmt"
)

func main() {
	euler()
}

func euler() {
	// euler
	// use complex
	c := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(c)
	fmt.Printf("%.3f/n", c)
}
