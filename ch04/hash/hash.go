package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable) {
	for _, node := range hash.Table {
		for node != nil {
			fmt.Printf("%d -> ", node.Value)
			node = node.Next
		}
		fmt.Println()
	}
}

func main() {
	table := make(map[int]*Node, 10)
	hash := &HashTable{Table: table, Size: 10}
	fmt.Println("Number of spaces:", hash.Size)
	for i := 0; i < 95; i++ {
		insert(hash, i)
	}
	traverse(hash)
}
