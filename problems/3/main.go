package main

import (
	"math"
	"math/big"
	"testing"
)

func a(n int) int {
	for i := int(math.Sqrt(float64(n))); i > 1; i-- {
		if n%i == 0 && big.NewInt(int64(i)).ProbablyPrime(20) {
			return i
		}
	}
	return 0
}

func main() {
	println("a: ", a(600851475143)) // 6857
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(600851475143)
		}
	})
	println("a: ", a.String(), a.MemString())
}
