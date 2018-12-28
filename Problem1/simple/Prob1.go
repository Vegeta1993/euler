package main

import (
	"fmt"
)

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
