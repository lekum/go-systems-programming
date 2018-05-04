package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Printf("Usage: %s <file1> [<file2> [... <fileN>]]\n",
			filepath.Base(args[0]))
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

			ip := findIP(line)
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
