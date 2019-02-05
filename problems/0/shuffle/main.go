package main

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func verify(perms [][]int) []float64 {
	return nil
}

func b(size int) (perm []int) {
	for i := 0; i < size; i++ {
		perm = append(perm, i)
	}
	rand.Shuffle(len(perm), func(i, j int) {
		perm[i], perm[j] = perm[j], perm[i]
	})
	return
}

func a(size int) (perm []int) {
	for i := 0; i < size; i++ {
		perm = append(perm, i)
	}
	n := len(perm)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		perm[r], perm[i] = perm[i], perm[r]
	}

	return
}

// c todo

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	log.Println("a:", a(10))
	log.Println("b:", b(10))
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(10)
		}
	})
	println("a: ", a.String(), a.MemString())

	b := testing.Benchmark(func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			b(10)
		}
	})
	println("b: ", b.String(), b.MemString())
}
