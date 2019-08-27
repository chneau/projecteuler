// https://projecteuler.net/problem=32
package main

import (
	"math"
	"testing"
)

// max = 9 => pandi from 1 to 9
func pandigital(test, max int) bool {
	bb := make([]byte, max)
	for i := test; i != 0; i /= 10 {
		o := (i % 10) - 1
		if o == -1 || o >= max {
			return false
		}
		if bb[o] == 1 {
			return false
		}
		bb[o] = 1
	}
	for i := 0; i < max; i++ {
		if bb[i] == 0 {
			return false
		}
	}
	return true
}

func concat(a, b int) int { // not idiot proof, no check if overflow (near above 1e18)
	for i := 10; ; i *= 10 {
		if b < i {
			return i*a + b
		}
	}
}

func a(max int) (sum int) {
	memory := map[int]struct{}{}
	tresh := int(math.Pow10(max))
	for i := 0; i < 1e4; i++ { // how to estime this 1e4 (got it from err and trials)
		for j := 0; j < 1e4; j++ {
			if j > i { // halv to 2.5s
				break
			}
			prod := i * j
			if prod > tresh {
				break
			}
			if pandigital(concat(concat(i, j), prod), max) {
				if _, exist := memory[prod]; exist {
					continue
				}
				memory[prod] = struct{}{}
				sum += i * j
			}
		}
	}
	return
}

func main() {
	res := a(9)
	println("a:", res) // 45228
	aaa := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			a(9) // 2.5s
		}
	})
	println("a:", aaa.String(), aaa.MemString())
}

// Solution is 45228
