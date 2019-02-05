//https://projecteuler.net/problem=21
package main

import (
	"math"
	"testing"
)

func sumOfDivisors(n int) (sum int) {
	sum = 1 // but we don't want n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sum += i
			sum += n / i
		}
	}
	return
}

func a(top int) (sum int) {
	xx := map[int]int{}
	for i := 1; i < top; i++ {
		xx[i] = sumOfDivisors(i)
	}
	for i := 1; i < top; i++ {
		b := xx[i]
		if b > top {
			continue
		}
		x := xx[b]
		if x == i && i != b {
			sum += x
		}

	}
	return
}

func b(top int) (sum int) {
	for i := 1; i < top; i++ {
		b := sumOfDivisors(i)
		if b > top {
			continue
		}
		x := sumOfDivisors(b)
		if x == i && i != b {
			sum += x
		}
	}
	return
}

func c(top int) (sum int) {
	xx := make([]int, top) // could use slice since we know we will use 1 to top
	// map is good only if we don't know how is the key distributed
	for i := 1; i < top; i++ {
		xx[i] = sumOfDivisors(i)
	}
	for i := 1; i < top; i++ {
		b := xx[i]
		if b > top {
			continue
		}
		x := xx[b]
		if x == i && i != b {
			sum += x
		}

	}
	return
}

func main() {
	res := a(10000)
	println("a:", res)
	res = b(10000)
	println("b:", res)
	res = c(10000)
	println("c:", res)
	aa := testing.Benchmark(func(b *testing.B) { // 60% faster than b but alloc
		for i := 0; i < b.N; i++ {
			a(10000)
		}
	})
	println("a:", aa.String(), aa.MemString())
	bb := testing.Benchmark(func(bb *testing.B) { // slow than a but no alloc
		for i := 0; i < bb.N; i++ {
			b(10000)
		}
	})
	println("b:", bb.String(), bb.MemString())
	cc := testing.Benchmark(func(bb *testing.B) { // wee (weeeeeeeee) faster than a, less allocs (one time alloc). good compromise between the two solutions
		for i := 0; i < bb.N; i++ {
			c(10000)
		}
	})
	println("c:", cc.String(), cc.MemString())
}

// Solution is 31626
