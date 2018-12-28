/*
goos: linux
goarch: amd64
pkg: euler/Problem3/suggested
BenchmarkSimple-8   	  200000	      8394 ns/op
PASS
ok  	euler/Problem3/suggested	1.784s

*/
package main

import (
	"testing"
)

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeFactors(600851475143)
	}
}
