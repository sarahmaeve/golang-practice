package main

// count words in a string
// from Golang tour

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordmap := make(map[string]int)
	for _, v := range words {
		wordmap[v] += 1
	}
	return wordmap
}

func main() {
	wc.Test(WordCount)
}
