package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	/*
		 A slice does not store any data, it just describes a section of an underlying array.

		Changing the elements of a slice modifies the corresponding elements of its underlying array.

		Other slices that share the same underlying array will see those changes.
	*/

	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
