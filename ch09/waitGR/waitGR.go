package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Waiting for goroutines!")
	var wg sync.WaitGroup

	const numGR = 100
	wg.Add(numGR)
	var i int64
	for i = 0; i < numGR; i++ {
		go func(x int64) {
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	wg.Wait()
	fmt.Println("\nExiting...")
}
