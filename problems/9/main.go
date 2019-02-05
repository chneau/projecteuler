package main

import (
	"testing"
)

func checkUnderstanding(n int) (a, b, c int) {
	max := n
	for a = 1; a < max; a++ {
		for b = a + 1; b < max; b++ {
			for c = b + 1; c <= max; c++ {
				if a*a+b*b == c*c && c*c == n {
					return
				}
			}
		}
	}
	return 0, 0, 0
}

func a(n int) (a, b, c int) { // quiet slow way of doing things
	max := n / 2 // not fully sure, x5 faster tho
	for a = 1; a < max; a++ {
		for b = a + 1; b < max; b++ {
			for c = b + 1; c <= max; c++ {
				if a*a+b*b == c*c && a+b+c == n {
					return
				}
			}
		}
	}
	return 0, 0, 0
}

func main() {
	if aa, bb, cc := checkUnderstanding(25); aa != 3 || bb != 4 || cc != 5 {
		println(aa, bb, cc)
		panic("aa, bb, cc := a(25); aa != 3 || bb != 4 || cc != 5")
	}
	aa, bb, cc := a(1000) // 200, 375, 425
	println("a:", aa, bb, cc)
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(1000)
		}
	})
	println("a: ", a.String(), a.MemString())
}
