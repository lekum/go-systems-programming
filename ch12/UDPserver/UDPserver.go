package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <port>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	addr, err := net.ResolveUDPAddr("udp", ":"+os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := c.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n]))
		data := []byte(buffer[0:n])
		_, err = c.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP server")
			return
		}
	}
}
