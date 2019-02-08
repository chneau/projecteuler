// https://projecteuler.net/problem=29
package main

import (
	"math"
	"strconv"
	"testing"
)

func p16(x, n int) (sum string) { // from problem 16, modified
	arr := []int{x}
	for i := 1; i < n; i++ {
		arrlen := len(arr)
		// does multiplication number by number
		for j := 0; j < arrlen; j++ {
			arr[j] *= x
		}
		// sort of shift rest like we are 4yo
		// eg: 7*7 = 49, 49*7 = 4*7(*10) + 9*7 = 280 + 63 = 343
		for j := arrlen - 1; j >= 0; j-- {
			if arr[j] >= 10 {
				if j >= 1 {
					arr[j-1] += arr[j] / 10
					arr[j] %= 10
				} else {
					new := arr[j] / 10
					arr[j] %= 10
					arr = append([]int{new}, arr...)
				}
			}
		}
		// flatten arr[0] in case x is >= 10
		// easier to understand if you log there and then
		// eg [105 9 8] => [1 0 5 9 8]
		for arr[0] >= 10 {
			new := arr[0] / 10
			arr[0] %= 10
			arr = append([]int{new}, arr...)
		}
	}
	// calculate the sum
	for i := 0; i < len(arr); i++ {
		sum += strconv.Itoa(arr[i])
	}
	return
}

func a(mina, maxa, minb, maxb float64) int { // total failure but almost there ... TODO: copy past the 4yo multiplicator
	tree := map[float64]struct{}{}
	for a := mina; a <= maxa; a++ {
		for b := minb; b <= maxb; b++ {
			tree[math.Pow(a, b)] = struct{}{}
		}
	}
	return len(tree)
}

func b(mina, maxa, minb, maxb int) int {
	tree := map[string]struct{}{}
	for a := mina; a <= maxa; a++ {
		for b := minb; b <= maxb; b++ {
			tree[p16(a, b)] = struct{}{}
		}
	}
	return len(tree)
}

func powint(a, b int) (res int) {
	res = 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return
}

func c(mina, maxa, minb, maxb int) int {
	tree := map[int]struct{}{}
	for a := mina; a <= maxa; a++ {
		for b := minb; b <= maxb; b++ {
			tree[powint(a, b)] = struct{}{}
		}
	}
	return len(tree)
}

func main() {
	res := a(2, 100, 2, 100)
	println("a:", res, "FAIL") // 9220
	res = b(2, 100, 2, 100)
	println("b:", res) // 9183
	res = c(2, 100, 2, 100)
	println("c:", res, "FAIL") // 6413
	aaa := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			b(2, 100, 2, 100) // slow.
		}
	})
	println("b:", aaa.String(), aaa.MemString())
}

// Solution is 9183
