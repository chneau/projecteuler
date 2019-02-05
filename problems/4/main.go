package main

import (
	"math"
	"strconv"
	"testing"
)

func a(n int) (xx int, yy int, r int) {
	max := int(math.Pow10(n)) - 1
	for x := max; x > 1; x-- {
		for y := max; y >= x; y-- {
			a := x * y
			if r > a {
				break
			}
			strA := strconv.Itoa(a)
			lenStrA := len(strA)
			isPalindrome := true
			for i := 0; i < lenStrA/2; i++ {
				if strA[i] == strA[lenStrA-1-i] {
					continue
				}
				isPalindrome = false
			}
			if isPalindrome {
				if a > r {
					r = a
					xx = x
					yy = y
				}
			}
		}
	}
	return xx, yy, r
}

func main() { // should be 906609
	x, y, r := a(3) // 913 993 906609
	println(x, y, r)
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(3)
		}
	})
	println("a: ", a.String(), a.MemString())
}
