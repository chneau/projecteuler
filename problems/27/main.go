// https://projecteuler.net/problem=27
package main

import (
	"math"
	"testing"
)

func isPrime(x int64) bool {
	if x < 1 {
		return false
	}
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

func a(maxa, maxb int64) (a, b int64) {
	best := int64(0)
	for ia := -maxa + 1; ia < maxa; ia++ {
		for ib := -maxb; ib <= maxb; ib++ {
			i := int64(0)
			for ; ; i++ {
				if !isPrime(i*i + ia*i + ib) {
					break
				}
			}
			if i > best {
				best = i
				a = ia
				b = ib
			}
		}
	}
	return
}

func main() {
	aa, bb := a(1000, 1000)
	println("a:", aa*bb)
	aaa := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(1000, 1000)
		}
	})
	println("a:", aaa.String(), aaa.MemString())
}

// Solution is -59231
