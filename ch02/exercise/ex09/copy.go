package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("copy(a, b)")
	n := copy(a, b)
	fmt.Printf("Copied %d elements\n", n)
	fmt.Println("a:", a)
	fmt.Println("copy(b, a)")
	n = copy(b, a)
	fmt.Printf("Copied %d elements\n", n)
	fmt.Println("b:", b)
}
