package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	r := make(map[string]int)

	fields := strings.Fields(s)

	for _, field := range fields {
		r[field] += 1
	}
	return r
}

func main() {
	wc.Test(WordCount)
}
