package main

import (
	"fmt"
	"testing"
)

func sign(w string) int {
	s := 10
	for i := 0; i+1 < len(w); i++ {
		currentLetter := int(w[i])
		nextLetter := int(w[i+1])
		s ^= nextLetter - currentLetter
	}
	return s
}

func a(ww []string) (bb [][]string) {
	allSigns := map[int][]string{}
	for _, w := range ww {
		s := sign(w)
		if _, exist := allSigns[s]; exist {
			allSigns[s] = append(allSigns[s], w)
			continue
		}
		allSigns[s] = []string{w}
	}
	for _, ww := range allSigns {
		if len(ww) > 1 {
			bb = append(bb, ww)
		}
	}
	return
}

func b(ww []string) (bb [][]string) {
	allSigns := map[int][]int{}
	for i, w := range ww {
		s := sign(w)
		if _, exist := allSigns[s]; exist {
			allSigns[s] = append(allSigns[s], i)
			continue
		}
		allSigns[s] = []int{i}
	}
	for _, ii := range allSigns {
		if len(ii) > 1 {
			ws := []string{}
			for _, i := range ii {
				ws = append(ws, ww[i])
			}
			bb = append(bb, ws)
		}
	}
	return
}

// backToA make every word start with letter A
func backToA(ww []string) (rr []string) {
	for _, w := range ww {
		offset := int(w[0]) - int('a')
		newW := ""
		for _, l := range w {
			newL := string(int(l) - offset)
			newW += newL
		}
		rr = append(rr, newW)
	}
	return
}

func c(ww []string) (bb [][]string) {
	rr := backToA(ww)
	bags := map[string][]string{}
	for _, r := range rr {
		if _, exist := bags[r]; exist {
			bags[r] = append(bags[r], r)
			continue
		}
		bags[r] = []string{r}
	}
	for _, b := range bags {
		if len(b) > 1 {
			bb = append(bb, b)
		}
	}
	return
}

func main() {
	words := []string{"abcd", "bcde", "dog", "god", "eph", "acdc", "bded", "a", "b", "c"}
	bags := a(words)
	fmt.Println("a:", bags)
	bags = b(words)
	fmt.Println("b:", bags)
	bags = c(words)
	fmt.Println("c:", bags)
	a := testing.Benchmark(func(b *testing.B) { // original idea
		for i := 0; i < b.N; i++ {
			a(words)
		}
	})
	println("a: ", a.String(), a.MemString())
	b := testing.Benchmark(func(bb *testing.B) { // Not this faster, use more memory ...
		for i := 0; i < bb.N; i++ {
			b(words)
		}
	})
	println("b: ", b.String(), b.MemString())
	c := testing.Benchmark(func(b *testing.B) { // slower but the idea is nice
		for i := 0; i < b.N; i++ {
			c(words)
		}
	})
	println("c: ", c.String(), c.MemString())
}
