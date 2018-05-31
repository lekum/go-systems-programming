package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := ""
	var f *os.File
	args := os.Args
	if len(args) == 1 {
		f = os.Stdin
	} else {
		filename = args[1]
		fileHandler, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening %s: %s", filename, err)
			os.Exit(1)
		}
		f = fileHandler
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
