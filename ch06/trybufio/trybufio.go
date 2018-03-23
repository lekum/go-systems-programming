package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if scanner.Err() != nil {
			fmt.Printf("Error reading file %v\n", err)
			os.Exit(1)
		}
		fmt.Println(line)
	}
}
