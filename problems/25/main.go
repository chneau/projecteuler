// https://projecteuler.net/problem=25
package main

func fibonacci(n int) (x int) {
	if n < 3 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func a(length int) int {
	n1 := make([]int, length, length) // is basically n-1
	n2 := make([]int, length, length) // and n-2
	a, b := n1, n2
	n1[0], n2[0] = 1, 1
	for i := 3; ; i++ { // and the idea is to add n-2 to n-1, which then swap both
		a = n1
		b = n2
		if i%2 != 0 {
			a = n2
			b = n1
		}
		for n := 0; n < length; n++ {
			if n > 0 && a[n-1] > 9 {
				a[n] += int(a[n-1] / 10)
				a[n-1] %= 10
			}
			a[n] += b[n]
		}
		if a[length-1] != 0 {
			return i
		}
	}
}

func main() {
	sol := a(1e3)
	println("a:", sol)
}

// Solution is 4782
