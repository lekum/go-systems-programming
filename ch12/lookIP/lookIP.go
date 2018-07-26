package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <ip>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	IP := os.Args[1]
	addr := net.ParseIP(IP)
	if addr == nil {
		fmt.Println("Not a valid IP address!")
		os.Exit(1)
	}

	hosts, err := net.LookupAddr(IP)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, hostname := range hosts {
		fmt.Println(hostname)
	}
}
