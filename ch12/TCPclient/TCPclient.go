package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <address>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	addr := os.Args[1]
	c, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer c.Close()

	stdioReader := bufio.NewReader(os.Stdin)
	messageReader := bufio.NewReader(c)

	for {
		fmt.Print(">> ")
		text, _ := stdioReader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		message, _ := messageReader.ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}

}
