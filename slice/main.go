package main

import "fmt"

func main() {
	// array
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// slice of arr
	s := arr[2:6]
	fmt.Println(s)

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	// change slice, it's array will change too
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	// array change, it's slice will change too
	fmt.Println(s2)

	fmt.Println("Reslice")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println("s1 = ",s1)
	fmt.Println("s2 = ",s2)
}

func updateSlice(s []int) {
	s[0] = 100
}
