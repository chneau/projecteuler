package main

import (
	"math"
	"testing"
)

// define tip: the substr that is being compressed on the string
// eg:
// string -> "abcdabc"
// tip -> "abc"
// compressed string -> "+d+"

func a(s string) (tip string, compressed string) { // but doesnt work with abcXabcYabcLwwwwwwwwwwwwwwwwwwwwwwwwwww
	middle := len(s) / 2
	tree := map[string]int{}
	for mid := 2; mid < middle; mid++ {
		for j := 0; j+mid < len(s); j++ {
			substr := s[j : j+mid]
			if v, exist := tree[substr]; exist {
				tree[substr] = v + 1 - len(substr)
				continue
			}
			tree[substr] = 1
		}
	}
	bestCompression := 0
	for k, v := range tree {
		if v < bestCompression {
			bestCompression = v
			tip = k
		}
	}
	for i := 0; i < len(s); {
		j := i + len(tip)
		if j <= len(s) {
			if s[i:j] == tip {
				compressed += "+"
				i += len(tip)
				continue
			}
		}
		compressed += s[i : i+1]
		i++
	}
	return
}

func b(s string) (tip string, compressed string) { // find optimal solution but is slower since it compute/compress all possible tip
	compress := func(s string, tip string) (result string) {
		for i := 0; i < len(s); {
			j := i + len(tip)
			if j <= len(s) {
				if s[i:j] == tip {
					result += "+"
					i += len(tip)
					continue
				}
			}
			result += s[i : i+1]
			i++
		}
		return
	}
	middle := len(s) / 2
	tree := map[string]struct{}{}
	for mid := 2; mid < middle; mid++ {
		for j := 0; j+mid < len(s); j++ {
			substr := s[j : j+mid]
			tree[substr] = struct{}{}
		}
	}
	bestSize := math.MaxInt64
	for k := range tree {
		c := compress(s, k)
		newSize := len(c) + len(k)
		if newSize >= len(s) {
			continue
		}
		if newSize < bestSize {
			bestSize = newSize
			tip = k
			compressed = c
		}
	}
	return
}

func main() {
	str := "abcXabcYabcZ"
	t, c := a(str)
	println("a:", t, c)
	t, c = b(str)
	println("b:", t, c)
	a := testing.Benchmark(func(b *testing.B) { // unreliable but fast
		for i := 0; i < b.N; i++ {
			a(str)
		}
	})
	println("a: ", a.String(), a.MemString())
	b := testing.Benchmark(func(bb *testing.B) { // reliable but slow
		for i := 0; i < bb.N; i++ {
			b(str)
		}
	})
	println("b: ", b.String(), b.MemString())
}
