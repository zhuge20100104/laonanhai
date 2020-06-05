package fib

import (
	"testing"
)

func benchmarkFib(b *testing.B, size int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fib(size)
	}
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

func BenchmarkFib100(b *testing.B) {
	benchmarkFib(b, 100)
}
