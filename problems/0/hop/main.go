package main

import (
	"log"
	"testing"
)

type hop struct {
	src string
	dst string
}

func example() []hop {
	hh := []hop{
		{"a", "b"},
		{"a", "e"},
		{"b", "c"},
		{"c", "d"},
		{"f", "g"},
		{"f", "k"},
		{"x", "y"},
	}
	return hh
}

func a(hops []hop) (res []string) {
	tree := map[string]int{}
	for _, hop := range hops {
		tree[hop.src] += 0
		tree[hop.dst]++
	}
	for k, v := range tree {
		if v == 0 {
			res = append(res, k)
		}
	}
	return
}

func main() {
	ex := example()
	log.Println("a:", a(ex))
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(ex)
		}
	})
	println("a: ", a.String(), a.MemString())
}
