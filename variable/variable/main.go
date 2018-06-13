package main

import "fmt"

// Declare and assign a variable out of func, you have yo use var keyword
// an out of func variable's scope is in package
var aa = 3
var ss = "kkk"
var bb = true

var (
	a1 = 3
	s1 = "kkk"
	b1 = true
)

func variableZeroValue() {
	// Declare a variable, it will be assigned a default value
	// A variable must be used in Go, or there will be a error
	var a int    // default value 0
	var s string // default value ""
	fmt.Printf("%d %q \n", a, s)
}

func variableInitialValue() {
	// Declare and initialize a variable
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// := can only use in a func
	a, b, c, s := 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("Hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
}