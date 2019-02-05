package main

import "testing"

func a(n int) (r int) {
	r = 1
	for ; ; r++ {
		good := true
		for i := 2; i <= n; i++ {
			if r%i != 0 {
				good = false
				break
			}
		}
		if good {
			return r
		}
	}
}

func main() {
	if a(10) != 2520 {
		panic("a(10) != 2520")
	}
	println("a: ", a(20)) // 232792560
	// TODO: too slow omg !
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(20)
		}
	})
	println("a: ", a.String(), a.MemString())
}
