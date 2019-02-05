package main

import "testing"

func a(n int) (sum int) {
	a := 1
	b := 1
	for b < n {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
		}
	}
	return
}

func main() {
	println("a: ", a(4e6)) // 4613732
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(4e6)
		}
	})
	println("a: ", a.String(), a.MemString())
}
