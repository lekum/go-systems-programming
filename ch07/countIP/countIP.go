package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	minusCOL := flag.Int("COL", 1, "Column")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Printf("Usage: %s <file1> [<file2> [... <fileN>]]\n",
			filepath.Base(args[0]))
		os.Exit(1)
	}

	column := *minusCOL
	if column < 0 {
		fmt.Println("Invalid column number!")
		os.Exit(1)
	}

	myIPs := make(map[string]int)

	for _, filename := range args {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error reading file: %s\n", err)
				continue
			}

			data := strings.Fields(line)
			ip := data[column-1]
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			_, ok := myIPs[ip]
			if ok {
				myIPs[ip] += 1
			} else {
				myIPs[ip] = 1
			}
		}
	}

	for key, val := range myIPs {
		fmt.Printf("%s %d\n", key, val)
	}
}
