/*
TITLE: Multiples of 3 and 5

DESCRIPTION: If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.

*/

package main

import (
	"fmt"
)

// Calculates sum of numbers divisible by 3 and 5

func Sum(n int) int {

	var s int

	for i := 3; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			s += i
		}
	}

	return s
}
func main() {

	// Variable *sum* consists of sum of all multiples of 3 and 5
	var sum int

	sum = Sum(1000)

	fmt.Println(sum)

}
