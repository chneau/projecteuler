package main

import "testing"

func collatz(x int) (seqSize int) {
	for x != 1 {
		if x%2 == 0 {
			x /= 2
		} else {
			x = 3*x + 1
		}
		seqSize++
	}
	seqSize++
	return
}

func a(before int) (best, size int) {
	for n := 1; n <= before; n++ {
		score := collatz(n)
		if score > size {
			size = score
			best = n
		}
	}
	return
}

func main() {
	best, _ := a(1e6)
	println("a:", best)
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(1e6)
		}
	})
	println("a: ", a.String(), a.MemString())
}

// Solution is 837799
