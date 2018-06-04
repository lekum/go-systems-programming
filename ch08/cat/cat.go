package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func catFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil

}

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		os.Exit(0)
	}
	files := os.Args[1:]
	for _, filename := range files {
		err := catFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}

	}
}
