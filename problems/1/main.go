package main

import (
	"testing"
)

func a(n int) (r int) { // faster
	for i := 0; i < n; i++ {
		if i%5 == 0 || i%3 == 0 {
			r += i
		}
	}
	return
}

func bb(n int) (r int) {
	d := map[int]struct{}{}
	for i := 3; i < n; i = i + 3 {
		d[i] = struct{}{}
	}
	for i := 5; i < n; i = i + 5 {
		d[i] = struct{}{}
	}
	for i := range d {
		r += i
	}
	return
}

func main() {
	println("a: ", a(1000))  // 233168
	println("b: ", bb(1000)) // 233168
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(1000)
		}
	})
	println("a: ", a.String(), a.MemString())
	b := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bb(1000)
		}
	})
	println("b: ", b.String(), b.MemString())
}
