// https://projecteuler.net/problem=23
package main

import (
	"math"
	"testing"

	"github.com/chneau/tt"
)

// LIMIT is the max by the theory
const LIMIT = 28123

// comming from problem 21
func sumOfDivisors(num int) (sum int) {
	sroot := int(math.Sqrt(float64(num))) + 1
	for i := 1; i < sroot; i++ {
		if num%i == 0 {
			if i*i != num {
				sum += i
				sum += num / i
			} else {
				sum += i
			}
		}
	}
	sum -= num
	return
}

func findAbundantNumbers() (abundants []int) { // ~ 38ms
	defer tt.T()()
	for i := 1; i < LIMIT; i++ {
		res := sumOfDivisors(i)
		if i < res {
			abundants = append(abundants, i)
		}
	}
	return
}

func a(abundantNumbers []int) (sum int) { // ~ 740ms
	tree := map[int]struct{}{}
	for _, x := range abundantNumbers {
		for _, y := range abundantNumbers {
			sum := x + y
			if sum > LIMIT {
				break
			}
			tree[sum] = struct{}{}
		}
	}
	for i := 0; i < LIMIT; i++ {
		if _, in := tree[i]; !in {
			sum += i
		}
	}
	return
}

func main() {
	ann := findAbundantNumbers()
	sum := a(ann)
	println("a:", sum)
	aa := testing.Benchmark(func(b *testing.B) { // not happy with this solution for the moment but ok
		for i := 0; i < b.N; i++ {
			a(ann)
		}
	})
	println("a:", aa.String(), aa.MemString())
}

// Solution is 4179871
