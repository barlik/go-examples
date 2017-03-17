package main

import (
	"fmt"
)

func main() {
	WORDS := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	if word, ok := WORDS["second"]; !ok {
		fmt.Println("Doesn't exist\n")
	} else {
		fmt.Printf("%d\n", word)
	}
}
