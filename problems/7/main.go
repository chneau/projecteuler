package main

import (
	"math"
	"math/big"
	"testing"
)

func isPrime(x int64) bool {
	if x == 2 {
		return true
	}
	z := int64(math.Sqrt(float64(x)))
	for i := int64(2); i <= z; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func a(n int) int64 {
	p := 0
	for i := int64(2); ; i++ {
		if big.NewInt(i).ProbablyPrime(0) { // <- pretty slow
			p++
			if p == n {
				return i
			}
		}
	}
}

func b(n int) int64 {
	p := 0
	for i := int64(2); ; i++ {
		if isPrime(i) {
			p++
			if p == n {
				return i
			}
		}
	}
}

func main() {
	if a(6) != 13 {
		panic("a(6) != 13")
	}
	println("a: ", a(10001)) // 104743
	println("b: ", b(10001)) // 104743
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(10001)
		}
	})
	println("a: ", a.String(), a.MemString())
	b := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			b(10001)
		}
	})
	println("b: ", b.String(), b.MemString())
}
