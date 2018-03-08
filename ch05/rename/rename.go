package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	minusOverrite := flag.Bool("overwrite", false, "overwrite")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Please provide two arguments!")
		os.Exit(1)
	}

	source := args[0]
	destination := args[1]
	fileInfo, err := os.Stat(source)
	if err != nil {
		fmt.Println("Error reading:", source, err)
		os.Exit(1)
	}
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println("Only regular files are supported as source")
		os.Exit(1)
	}

	newDestination := destination
	destInfo, err := os.Stat(destination)
	if err == nil {
		mode = destInfo.Mode()
		if mode.IsDir() {
			justTheName := filepath.Base(source)
			newDestination = destination + "/" + justTheName
		}
	}
	destination = newDestination
	destInfo, err = os.Stat(destination)
	if err == nil {
		if !*minusOverrite {
			fmt.Println("Destination file already exists!")
			os.Exit(1)
		}
	}
	err = os.Rename(source, destination)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
