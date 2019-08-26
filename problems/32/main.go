// https://projecteuler.net/problem=32
package main

import "log"

func concat(a, b int) int { // not idiot proof, no check if overflow (near above 1e18)
	for i := 10; ; i *= 10 {
		if b < i {
			return i*a + b
		}
	}
}

func main() {
	log.Println(concat(1e10, 1e7))
}

// Solution is 45228
