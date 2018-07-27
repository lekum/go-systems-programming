package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <hostname>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	hostname := os.Args[1]

	ips, err := net.LookupHost(hostname)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
