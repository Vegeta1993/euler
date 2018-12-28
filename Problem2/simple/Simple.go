package main

import (
	"fmt"
)

func Fib(a, b int) (int, int) {
	return b, a + b
}

func EvenSum(end int) int {

	a := 1
	b := 1

	var sum int

	for b < end {
		if b%2 == 0 {
			sum += b
		}
		a, b = Fib(a, b)
	}
	return sum
}

func main() {
	fmt.Println(EvenSum(4000000))
}
