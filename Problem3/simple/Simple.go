/*

TITLE: Largest prime factor

DESCRIPTION: The prime factors of 13195 are 5, 7, 13 and 29. What is the largest prime factor of the number 600851475143?

*/

package main

import (
	"fmt"
)

func PrimeFactors(num int) int {
	/*

	    Here num is checked against all i. Considering the fact that only 2 is prime even number, we need not check for 4,6 and so on.

	   And second optimisation is that we check against the number itself, while we know for a fact that for any number we can only have atmost one prime factor greater than sqrt(n) (if n is prime number). So we can create a maxfactor of sqrt(n) and limit i to this maxfactor.

	*/
	var largestPrime int

	for i := 2; i < num; i++ {
		if num%i == 0 {
			largestPrime = num / i
			for j := 2; j < i/2; j++ {
				if i%j != 0 {
					if largestPrime < i {
						largestPrime = i
					}
				}
			}
			num /= i
		}
	}

	return largestPrime
}

func main() {
	fmt.Println(PrimeFactors(600851475143))
}
