package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <port>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	addr := ":" + os.Args[1]

	s, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer l.Close()

	buffer := make([]byte, 1024)

	for {
		conn, err := l.Accept()
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Print("> ", string(buffer[0:n]))

		_, err = conn.Write(buffer)

		conn.Close()
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}

	}

}
