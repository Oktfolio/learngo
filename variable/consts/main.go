package main

import (
	"math"
	"fmt"
)

func main() {
	consts()
	enums()
	enumsWithIota()
}

func consts() {
	// a const's type is uncertain
	// you are allowed to name a const with lowercase
	// cause uppercase initial means public in Go
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	// enum in Go
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
}

func enumsWithIota() {
	const (
		cpp        = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, python, golang, javascript)

	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
