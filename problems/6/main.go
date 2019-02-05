package main

import "testing"

func sumFirst(n int) (r int) {
	for i := 1; i <= n; i++ {
		r += i * i
	}
	return
}

func squareFirst(n int) (r int) {
	for i := 1; i <= n; i++ {
		r += i
	}
	r *= r
	return
}

func a(n int) int {
	sq := 0
	su := 0
	for i := 1; i <= n; i++ { // 2 times faster if both += in same loop (120ns/op -> 70ns/op)
		// see c
		sq += i
		su += i * i
	}
	sq *= sq
	return sq - su
}

func b(n int) int {
	sq := n * (n + 1) / 2
	sq *= sq
	su := 0
	for i := 1; i <= n; i++ {
		su += i * i
	}
	return sq - su
}

func c(n int) int {
	sq := 0
	for i := 1; i <= n; i++ {
		sq += i
	}
	su := 0
	for i := 1; i <= n; i++ {
		su += i * i
	}
	sq *= sq
	return sq - su
}

func main() {
	if a(10) != 2640 {
		panic("a(10) != 2640")
	}
	println("a: ", a(100)) // 25164150
	println("b: ", b(100)) // 25164150
	println("c: ", c(100)) // 25164150
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(100)
		}
	})
	println("a: ", a.String(), a.MemString())
	b := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			b(100)
		}
	})
	println("b: ", b.String(), b.MemString())
	c := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c(100)
		}
	})
	println("c: ", c.String(), c.MemString())
}
