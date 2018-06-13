package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func main() {
	fmt.Println(eval(3, 4, "*"))
	q, r := div(13, 3)
	fmt.Println(q, r)
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(i int, i2 int) int {
		return int(math.Pow(float64(i), float64(i2)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsuppoerted operation: %s", op)
	}
}

// not suggested
/*func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}*/

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args (%d %d) ", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}
