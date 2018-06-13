package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello Go 语言" // UTF-8
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println(len(s))

	// ch is a rune
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", r)
	}

	fmt.Println()

	for i, r := range []rune(s) {
		fmt.Printf("%d %c; ", i, r)
	}
	fmt.Println()
}
