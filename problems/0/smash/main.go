package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var dictionary map[string]bool

func init() {
	dictionary = readWords() // for benchmark sake and simpleness of this file
}

func readWords() map[string]bool {
	words := map[string]bool{}
	b, err := ioutil.ReadFile("problems/0/smash/words_alpha.txt")
	if err != nil {
		panic(err)
	}
	s := string(b)
	rawWords := strings.Split(s, "\r\n")
	for i := range rawWords {
		words[rawWords[i]] = false
	}
	return words
}

func a(word string) bool {
	_, exist := dictionary[word]
	if !exist {
		return false
	}

	w := word[:]
	for len(w) > 1 {
		for i := 0; i < len(w); i++ {
			newWord := ""
			for j := 0; j < len(w); j++ {
				if j != i {
					newWord += w[j : j+1]
				}
			}
			_, exist := dictionary[newWord]
			if exist {
				w = newWord
				continue
			}
			return false
		}
	}
	return true
}

func longestSmashableWord() (longest string) {
	for word := range dictionary {
		if a(word) {
			if len(word) > len(longest) {
				longest = word
			}
		}
	}
	return
}

func main() {
	word := longestSmashableWord()
	fmt.Println("a:", a(word))
	a := testing.Benchmark(func(b *testing.B) { // original idea
		for i := 0; i < b.N; i++ {
			a(word)
		}
	})
	println("a: ", a.String(), a.MemString())
}
