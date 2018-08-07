package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <address>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	addr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	data := []byte("Hello UDP echo server!\n")
	_, err = c.Write(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Reply: ", string(buffer[0:n]))
}
