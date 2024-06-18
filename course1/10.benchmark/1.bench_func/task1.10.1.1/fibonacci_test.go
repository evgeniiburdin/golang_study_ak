package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}
