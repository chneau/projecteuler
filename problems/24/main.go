// https://projecteuler.net/problem=24
package main

import (
	"fmt"
	"log"
)

func factorial(n int) (sum int) {
	sum = 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return
}

func a(perm []string, nth int) (sol []string) {
	x := factorial(len(perm))
	log.Println(x)
	copy := append(perm[:0:0], perm...)
	log.Println(copy)
	for len(copy) > 0 {
		i := nth % len(perm)
		sol = append(sol, copy[i])
		copy = append(copy[:i], copy[i+1:]...)
		log.Println(sol)
		nth /= len(copy)
	}
	return perm
}

func main() {
	res := a([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, 1e6)
	fmt.Println("a:", res)
}

// Solution is 2783915460
