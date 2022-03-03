package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func evel(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func sum(num ...int) int {
	s := 0
	for i := range num {
		s += num[i]
	}
	return s
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+""+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	if result, err := evel(3, 4, "x"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(evel(2, 4, "*"))
	fmt.Println(div(13, 4))
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(sum(1, 2, 4, 3, 5))
}
