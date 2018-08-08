package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <port>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	addr := ":" + os.Args[1]
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {

	defer c.Close()

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print("-> ", string(netData))
		c.Write([]byte(netData))
		if strings.TrimSpace(string(netData)) == "STOP" {
			break
		}
	}
	time.Sleep(3 * time.Second)

}
