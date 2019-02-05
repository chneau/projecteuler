//https://projecteuler.net/problem=20
package main

import (
	"testing"
)

func factorial(n int) (arr []int) { // natural way of calculating stuff like a 4yo
	arr = append(arr, n)
	for i := n - 1; i > 0; i-- {
		for j := range arr {
			arr[j] *= i
		}
		arrlen := len(arr)
		for j := arrlen - 1; j >= 0; j-- {
			if arr[j] >= 10 {
				if j >= 1 {
					arr[j-1] += arr[j] / 10
					arr[j] %= 10
				} else {
					new := arr[j] / 10
					arr[j] %= 10
					arr = append([]int{new}, arr...) // this is alloc hungry
				}
			}
		}
		for arr[0] >= 10 {
			new := arr[0] / 10
			arr[0] %= 10
			arr = append([]int{new}, arr...)
		}
	}
	return
}

func reversefactorial(n int) (arr []int) { // keep array reversed, less allocs
	arr = append(arr, n)
	for i := n - 1; i > 0; i-- {
		for j := range arr {
			arr[j] *= i
		}
		arrlen := len(arr)
		for j := range arr {
			if arr[j] >= 10 {
				if j < arrlen-1 {
					arr[j+1] += arr[j] / 10
					arr[j] %= 10
				} else {
					new := arr[j] / 10
					arr[j] %= 10
					arr = append(arr, new) // this is alloc cool
				}
			}
		}
		for arr[len(arr)-1] >= 10 {
			new := arr[0] / 10
			arr[0] %= 10
			arr = append(arr, new)
		}

	}
	return
}

func a(n int) (sum int) {
	for _, v := range factorial(n) {
		sum += v
	}
	return
}
func b(n int) (sum int) {
	for _, v := range reversefactorial(n) {
		sum += v
	}
	return
}

func main() {
	res := a(100)
	println("a:", res)
	res = b(100)
	println("b:", res)
	aa := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(100)
		}
	})
	println("a: ", aa.String(), aa.MemString()) // 2x slower
	bb := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			b(100)
		}
	})
	println("b: ", bb.String(), bb.MemString()) // best atm
	// way of thinking:
	// 1. do something a 4yo do then
	// 2. optimise it (aka remove human weird way of doing things)
	// 3. ...
	// 4. profit
}

// Solution is 648
