package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"io"
	"strings"
)

func main() {
	fmt.Println(
		convertToBinary(5),       // 101
		convertToBinary(13),      // 1101
		convertToBinary(1548915), // 1101
		convertToBinary(0),       // 1101
	)
	printFile("abc.txt")
	s := `abc"Fds
		dsagdsaer
		gdsqre
		bte7

		32vfa"`
	printFileContents(strings.NewReader(s))
	// endlessLoop()
}

func convertToBinary(n int) string {
	result := ""
	if n == 0 {
		result = strconv.Itoa(0) + result
	} else {
		for ; n > 0; n /= 2 {
			lsb := n % 2
			result = strconv.Itoa(lsb) + result
		}
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func endlessLoop() {
	for {
		fmt.Println("abc")
	}
}