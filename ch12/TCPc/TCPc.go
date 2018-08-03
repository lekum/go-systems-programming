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

	addr := os.Args[1]
	myMessage := "Hello from TCP client!\n"

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()

	_, err = c.Write([]byte(myMessage))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print("-> ", myMessage)

	buffer := make([]byte, 1024)

	n, err := c.Read(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(">> ", string(buffer[0:n]))
}
