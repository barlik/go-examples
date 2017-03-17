package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var WORDS = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
}

func main() {
	fmt.Println(WORDS[2])
}
