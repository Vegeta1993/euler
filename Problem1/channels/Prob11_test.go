package main

import (
	"testing"
)

func benchmarksum(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(input)
	}
}

func BenchmarkThousand(b *testing.B)        { benchmarksum(1000, b) }
func BenchmarkMillion(b *testing.B)         { benchmarksum(1000000, b) }
func BenchmarkHundredThousand(b *testing.B) { benchmarksum(100000, b) }
