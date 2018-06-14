package main

import (
	"fmt"
	"os"
	"bufio"
	"learngo/functional/fib"
	"errors"
)

func tryDefer() {

	// defer First in Last out
	// defer guarantees that line will execute before return,
	// no matter whether there is any panic or some error or return,
	// you are supposed to defer when you open some resources which are supposed to close after using
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)


	err = errors.New("this is a custom err")
	if err != nil {
		fmt.Println("Error : ", err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	writeFile("fib.txt")
}
