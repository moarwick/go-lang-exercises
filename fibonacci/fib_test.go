package main

import "testing"

// run all benchmarks with: go test -bench=.
func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
