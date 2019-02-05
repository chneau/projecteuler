package main

import (
	"math"
	"runtime"
	"sync"
	"testing"

	"github.com/chneau/limiter"
)

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func a(n int) (s int) {
	for i := 0; i < n; i++ {
		if isPrime(i) {
			s += i
		}
	}
	return
}

func b(n int) (s int) {
	limit := limiter.New(runtime.NumCPU())
	m := sync.Mutex{}
	for i := 0; i < n; i++ {
		i := i
		limit.Execute(func() {
			if isPrime(i) {
				m.Lock()
				s += i
				m.Unlock()
			}
		})
	}
	limit.Wait()
	return
}

func c(n int) (s int) {
	isPrime := map[int]bool{}
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] == true {
			s += i
			for j := i; j <= n; j = j + i {
				isPrime[j] = false
			}
		}
	}
	return
}

func main() {
	if s := a(10); s != 17 {
		panic("s := a(10); s != 17")
	}
	// notes: a takes 9s for 2e5 ...
	// notes: b takes 2.4s for 2e5 ... (4 cores)
	// notes: c takes 50ms for 2e5 ... thanks to @jargnar https://github.com/jargnar/projecteuler
	// omiting a and b since slow
	println("a: too slow")
	println("b: too slow")
	s := c(2e6)
	println("c: ", s)
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(2e4)
		}
	})
	println("a: ", a.String(), a.MemString())
	b := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			b(2e4)
		}
	})
	println("b: ", b.String(), b.MemString())
	c := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c(2e4)
		}
	})
	println("c: ", c.String(), c.MemString())
}
