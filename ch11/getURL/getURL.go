package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	URL := os.Args[1]
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Printf("Error retrieving URL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Printf("Error copying content: %v\n", err)
		os.Exit(1)
	}
}
