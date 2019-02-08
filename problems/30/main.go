// https://projecteuler.net/problem=30
package main

import (
	"math"
	"testing"
)

func a(pow int) (sum int) {
	for i := 2; i < (pow-1)*int(math.Pow(9, float64(pow))); i++ {
		n := i
		nsum := 0
		for n != 0 {
			x := n % 10
			n /= 10
			nsum += int(math.Pow(float64(x), float64(pow)))
		}
		if i == nsum {
			sum += i
		}
	}
	return
}

func main() {
	res := a(5)
	println("a:", res) // 443839
	aaa := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			a(5) // 56ms
		}
	})
	println("a:", aaa.String(), aaa.MemString())
}

// Solution is 443839
