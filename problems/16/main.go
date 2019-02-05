//https://projecteuler.net/problem=16
package main

import (
	"testing"
)

func a(x, n int) (sum int) {
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
		sum += arr[i]
	}
	return
}

// solution inspired from github.com/swook
func b(pow int) (sum int) {
	n := make([]int, pow)
	n[0] = 1
	for i := 0; i < pow; i++ {
		for ii := 0; ii < len(n); ii++ {
			if n[ii] > 0 {
				n[ii] = n[ii] << 1
			}
			if ii > 0 && n[ii-1] >= 10 {
				n[ii] += int(n[ii-1] / 10)
				n[ii-1] = n[ii-1] % 10
			}
		}
	}
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return
}

func main() {
	res := a(2, 1000)
	println("a:", res)
	res = b(1000)
	println("b:", res)
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(2, 1000)
		}
	})
	println("a: ", a.String(), a.MemString())
	b := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			b(1000)
		}
	})
	println("b: ", b.String(), b.MemString())
	// a is faster but does more memory manipulation
	// a takes 2 parameters, which is a +
	// a must work with any positive x so far
	// b is bit slower but does less memory manipulation
}

// Solution is 1366
