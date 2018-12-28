package main

import (
	"fmt"
)

func Fib(a, b int) (int, int) {
	return b, (4*b + a)
}

func EvenSum(end int) int {

	a := 2
	b := 8

	var sum int = a

	for b < end {
		sum += b
		a, b = Fib(a, b)
	}
	return sum
}

func main() {
	fmt.Println(EvenSum(4000000))
}
