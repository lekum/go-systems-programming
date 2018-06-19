package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func searchStringInFile(s string, filename string, counts chan<- int, wg *sync.WaitGroup) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	occurrences := 0

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading line: %s\n", err)
			continue
		}

		n := strings.Count(line, s)
		occurrences += n
	}
	counts <- occurrences
	wg.Done()
	if occurrences > 0 {
		fmt.Printf("%d occurrences in %s\n", occurrences, filename)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <word> <file1> [<file2> [... <fileN>]]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	s := os.Args[1]
	counts := make(chan int)
	var wg sync.WaitGroup

	for _, f := range os.Args[2:] {
		wg.Add(1)
		go searchStringInFile(s, f, counts, &wg)
	}

	go func() {
		fmt.Println("Waiting")
		wg.Wait()
		fmt.Println("Finished Waiting")
		close(counts)
	}()

	count := 0
	for c := range counts {
		count += c
	}
	fmt.Printf("Total: %d occurrences\n", count)
}
