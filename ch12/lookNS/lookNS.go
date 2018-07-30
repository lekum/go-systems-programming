package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <domain>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	domain := os.Args[1]

	ns, err := net.LookupNS(domain)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, ns := range ns {
		fmt.Println(ns.Host)
	}
}
