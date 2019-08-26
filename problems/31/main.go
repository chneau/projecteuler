// https://projecteuler.net/problem=31
package main

import (
	"testing"
)

func a(target int) (count int) { // bruteforce
	for a := target; a >= 0; a -= 200 {
		for b := a; b >= 0; b -= 100 {
			for c := b; c >= 0; c -= 50 {
				for d := c; d >= 0; d -= 20 {
					for e := d; e >= 0; e -= 10 {
						for f := e; f >= 0; f -= 5 {
							for g := f; g >= 0; g -= 2 {
								count++
							}
						}
					}
				}
			}
		}
	}
	return
}

func b(target int) int { // source: https://www.mathblog.dk/project-euler-31-combinations-english-currency-denominations/
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	ways := make([]int, target+1)
	ways[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= target; j++ {
			ways[j] += ways[j-coins[i]]
		}
	}
	return ways[target]
}

func main() {
	res := a(200)
	println("a:", res) // 73682
	res = b(200)
	println("b:", res) // 73682
	aaa := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			a(200) // 41us
		}
	})
	println("a:", aaa.String(), aaa.MemString())
	bbb := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			b(200) // 1.6us
		}
	})
	println("b:", bbb.String(), bbb.MemString())
}

// Solution is 73682
