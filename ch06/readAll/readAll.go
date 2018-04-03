package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two command line arguments")
		os.Exit(1)
	}
	sourceFile := os.Args[1]
	destinationFile := os.Args[2]

	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Printf("Error creating the new file %s: %v\n", destinationFile, err)
		os.Exit(1)
	}
}
