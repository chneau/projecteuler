// https://projecteuler.net/problem=24
package main

import (
	"log"
	"testing"
)

func factorial(n int) (sum int) {
	sum = 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return
}

func a(perm []string, nth int) (solution []string) {
	nth--                               // since 1st is 1
	copy := append(perm[:0:0], perm...) // create a copy we items remove from
	for len(copy) > 0 {                 // while copy has items
		facto := factorial(len(copy) - 1)              // factorial of size-1
		index := nth / facto                           // calculate the index of the item to remove
		index %= len(copy)                             // modulo size to prevent overflow
		solution = append(solution, copy[index])       // add to solution and
		copy = append(copy[:index], copy[index+1:]...) // remove from copy
	}
	return
}

func main() {
	res := a([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, 1e6)
	log.Println("a:", res)
	aa := testing.Benchmark(func(b *testing.B) { // ~ 1179 ns/op. pretty good
		for i := 0; i < b.N; i++ {
			a([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, 1e6)
		}
	})
	println("a:", aa.String(), aa.MemString())
}

// feels: this one was interesting :)

// Solution is 2783915460
