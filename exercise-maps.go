package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")

	counter := make(map[string]int)
	for _, word := range words {
		counter[word] += 1
	}

	return counter
}

func main() {
	wc.Test(WordCount)
}
