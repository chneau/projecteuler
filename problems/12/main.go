package main

import (
	"log"
	"math"
	"testing"
)

func divisors(n int) (nb int) {
	if n == 1 {
		return 1
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 1; i < sqrt; i++ {
		if n%i == 0 {
			nb++
			nb++
		}
	}
	return
}

func a(n int) (index, result int) {
	if n%2 == 1 && n != 1 {
		return 0, 0
	}
	index = 1
	for {
		result += index
		d := divisors(result)
		if d >= n { // to have over
			break
		}
		index++
	}
	return
}

func main() {
	_, res := a(500)
	log.Println("a:", res)
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(500)
		}
	})
	println("a: ", a.String(), a.MemString())
}
