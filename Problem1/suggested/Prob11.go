/*
TITLE:Multiples of 3 and 5:

DESCRIPTION:If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.

*/
package main

import (
	"fmt"
)

/*

Calculates multiple of n which are less than m

*/
func Multiple(n, m int) int {
	var sum, mul int

	for i := 1; mul < m-n; i++ {
		mul = n * i
		sum += mul
	}

	return sum
}

// Add all multiples of 3 and 5 and substract common multiples (15)

func Sum(n int) int {

	return Multiple(3, n) + Multiple(5, n) - Multiple(15, n)

}
func main() {

	// find all multiples of 3 and 5 and deduct 15 from it
	var sum int

	sum = Sum(1000)

	fmt.Println(sum)

}
