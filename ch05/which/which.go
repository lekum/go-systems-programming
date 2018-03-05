package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	minusA := flag.Bool("a", false, "a")
	minusS := flag.Bool("s", false, "s")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}

	file := args[0]
	foundIt := false

	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")
	for _, directory := range pathSlice {
		fullPath := directory + "/" + file

		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					foundIt = true
					if *minusS == true {
						os.Exit(0)
					}
					if *minusA == true {
						fmt.Println(fullPath)
					} else {
						fmt.Println(fullPath)
						os.Exit(0)
					}
				}
			}
		}
	}
	if foundIt == false {
		os.Exit(1)
	}
}
