package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	ifCondition()
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(90),
		grade(100),
	)
}

func ifCondition() {
	const filename = "abc.txt"

	// bytes, err only work in if block
	if bytes, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", bytes)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		// panic will terminate the program
		panic(fmt.Sprintf("Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}