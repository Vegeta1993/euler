/*

TITLE: Largest prime factor

DESCRIPTION: The prime factors of 13195 are 5, 7, 13 and 29. What is the largest prime factor of the number 600851475143?

*/

package main

import (
	"fmt"
	"math"
)

func PrimeFactors(n int) int {
	var lastFactor int
	var factor int
	var maxFactor int

	if n%2 == 0 {
		lastFactor = 2
		n /= 2
		for n%2 == 0 {
			n /= 2
		}
	} else {
		lastFactor = 1
	}

	factor = 3
	maxFactor = int(math.Sqrt(float64(n)))
	for n > 1 && factor <= maxFactor {
		if n%factor == 0 {
			n /= factor
			lastFactor = factor
			for n%factor == 0 {
				n /= factor
			}
			maxFactor = int(math.Sqrt(float64(n)))
		}
		factor = factor + 2
	}

	if n == 1 {
		return lastFactor
	} else {
		return n
	}
}

func main() {
	fmt.Println(PrimeFactors(600851475143))
}
