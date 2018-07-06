package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s number\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	numGR, _ := strconv.ParseInt(os.Args[1], 10, 64)
	var wg sync.WaitGroup
	var i int64

	for i = 0; i < numGR; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%d ", i)
		}()
	}
	wg.Wait()
	fmt.Println("\nExiting...")
}
