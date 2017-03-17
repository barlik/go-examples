package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	for x := range a {
		for y := range a[x] {
			a[y][x] = uint8((y * x) % 100)
		}
	}

	return a
}

func main() {
	pic.Show(Pic)
}
