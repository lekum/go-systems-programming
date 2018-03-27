package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func countChars(r io.Reader) int {
	buf := make([]byte, 16)
	total := 0
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0
		}
		if err == io.EOF {
			break
		}
		total = total + n
	}
	return total
}

func writeNumberOfChars(w io.Writer, x int) {
	fmt.Fprintf(w, "%d\n", x)
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	_, err := os.Stat(filename)

	if err != nil {
		fmt.Printf("Error on file %s: %v\n", filename, err)
		os.Exit(1)
	}

	f, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	countFilename := strings.Join([]string{filename, "count"}, ".")
	f2, err := os.Create(countFilename)

	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", countFilename, err)
		os.Exit(1)
	}
	defer f2.Close()

	n := countChars(f)
	writeNumberOfChars(f2, n)
}
