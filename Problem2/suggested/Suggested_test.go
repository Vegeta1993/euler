package main

import (
	"testing"
)

func benchmarkSimple(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvenSum(input)
	}
}

func BenchmarkSimple4mil(b *testing.B)  { benchmarkSimple(4000000, b) }
func BenchmarkSimple40mil(b *testing.B) { benchmarkSimple(40000000, b) }
func BenchmarkSimple4bil(b *testing.B)  { benchmarkSimple(400000000, b) }
