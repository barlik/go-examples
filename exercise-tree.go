package main

import "fmt"
import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	walk(t, ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}

	ch <- t.Value
	// fmt.Printf("%d\n", t.Value)

	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	a := make(chan int)
	b := make(chan int)
	go Walk(t1, a)
	go Walk(t2, b)

	for value := range a {
		fmt.Println(value)
		value2, ok := <-b
		fmt.Println(value2)

		if value != value2 || !ok {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
