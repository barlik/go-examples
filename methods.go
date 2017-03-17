package main

import (
	"fmt"
	"math"
)

type Person struct {
	name, surname string
	age           int
}

func (person Person) ShowAge() {
	fmt.Println(person.age)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	me := Person{name: "Rastislav", surname: "Barlik", age: 30}
	me.ShowAge()
}
