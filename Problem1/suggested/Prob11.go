package main

import (
	"fmt"
)

func Multiple(n, m int) int {
	var sum, mul int

	for i := 1; mul < m-n; i++ {
		mul = n * i
		sum += mul
	}

	return sum
}

func Sum(n int) int {

	return Multiple(3, n) + Multiple(5, n) - Multiple(15, n)

}
func main() {

	// find all multiples of 3 and 5 and deduct 15 from it
	var sum int

	sum = Sum(1000)

	fmt.Println(sum)

}
