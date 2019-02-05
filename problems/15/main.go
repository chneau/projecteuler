//https://projecteuler.net/problem=15
package main

import (
	"sync"
	"testing"
)

// https://codereview.stackexchange.com/a/128838
func a(n int) (ret int) { // the mathematical trick
	ret = 1
	for i := 1; i <= n; i++ {
		ret *= n + i
		ret /= i
	}
	return
}

func b(right, down int) int { // sounds good, fork you
	if right == 0 || down == 0 {
		return 1
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	x := 0
	go func() {
		x = b(right-1, down)
		wg.Done()
	}()
	y := 0
	go func() {
		y = b(right, down-1)
		wg.Done()
	}()
	wg.Wait()
	return x + y
}

func c(right, down int) int { // slow
	if right == 0 || down == 0 {
		return 1
	}
	return c(right-1, down) + c(right, down-1)
}

// https://en.wikipedia.org/wiki/Lattice_path
func main() {
	// log.Println(b(20, 20)) // = fork bomb yourself.
	res := a(20)
	println("a:", res)
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(20)
		}
	})
	println("a: ", a.String(), a.MemString())
	// b and c are too slow
}

// Solution is 137846528820
