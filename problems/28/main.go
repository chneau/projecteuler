// https://projecteuler.net/problem=28
package main

import "testing"

func a(size int) (sum int) {
	sum = 1
	for i := 3; i <= size; i += 2 {
		// sum += i * i // top right corner
		// sum += i*i - (i - 1) // top left corner
		// sum += i*i - 2*(i-1) // bottom left corner
		// sum += i*i - 3*(i-1) // bottom right corner
		sum += 4*i*i - 6*(i-1) // this is faster
	}
	return
}

func b(size int) (sum int) {
	sum = 1
	for i := 3; i <= size; i += 2 {
		sum += 4*i*i - 6*i + 6 // this is slower
	}
	return
}

func c(size int) (sum int) {
	sum = 1
	for i := 3; i <= size; i += 2 {
		sum += 4*i*i - (6*i - 6) // this is as fast as a ... even a wee bit faster (5ns ~ 1%)
	}
	return
}

func main() {
	sol := a(1001)
	println("a:", sol)
	sol = b(1001)
	println("b:", sol)
	sol = c(1001)
	println("c:", sol)
	aaa := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(1001)
		}
	})
	println("a:", aaa.String(), aaa.MemString())
	aaa = testing.Benchmark(func(bb *testing.B) { // weirdly, this is a bit slower (501ns vs 435ns)
		for i := 0; i < bb.N; i++ {
			b(1001)
		}
	})
	println("b:", aaa.String(), aaa.MemString())
	aaa = testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			c(1001)
		}
	})
	println("c:", aaa.String(), aaa.MemString())
}

// Solution is 669171001
