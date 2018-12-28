/*
goos: linux
goarch: amd64
pkg: euler/Problem3/simple
BenchmarkSimple-8   	   20000	     87915 ns/op
PASS
ok  	euler/Problem3/simple	2.642s

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
