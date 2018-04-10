package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <FILENAME>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	in := string(buf)
	s := bufio.NewScanner(strings.NewReader(in))
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		fmt.Print(s.Text())
	}

}
