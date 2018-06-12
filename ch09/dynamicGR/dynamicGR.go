package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <num_goroutines>\n", filepath.Base(args[0]))
		os.Exit(1)
	}
	ng, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad goroutine number: %s\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(ng)
	for i := 0; i < ng; i++ {
		go func(x int) {
			defer wg.Done()
			fmt.Printf(" %d", x)
		}(i)
	}
	wg.Wait()
	fmt.Println("\nExiting...")
}
