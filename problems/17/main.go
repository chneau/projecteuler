//https://projecteuler.net/problem=17
package main

import (
	"strings"
	"testing"

	"github.com/divan/num2words" // too damn lazy for this one... thank you @divan
)

func a(n, m int) (sum int) {
	for i := n; i <= m; i++ {
		str := num2words.ConvertAnd(i)
		str = strings.Replace(str, " ", "", -1)
		str = strings.Replace(str, "-", "", -1)
		sum += len(str)
	}
	return
}

func main() {
	res := a(1, 1000)
	println("a:", res)
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(1, 1000)
		}
	})
	println("a: ", a.String(), a.MemString())
}

// Solution is 21124
