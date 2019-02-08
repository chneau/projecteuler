// https://projecteuler.net/problem=26
package main

import "testing"

func a(top int) (best, length int) {
	for i := top; i > 0; i-- {
		if length >= i {
			break
		}
		foundRemainders := map[int]struct{}{}
		value := 1
		for {
			remainder := value % i
			if _, exist := foundRemainders[remainder]; !exist {
				foundRemainders[remainder] = struct{}{}
				value = remainder * 10
			} else {
				break
			}
		}
		if len(foundRemainders) > length {
			length = len(foundRemainders)
			best = i
		}
	}
	return
}

func main() {
	best, _ := a(1e3)
	println("a:", best)
	aa := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(1e3)
		}
	})
	println("a:", aa.String(), aa.MemString())
}

// Solution is 983
