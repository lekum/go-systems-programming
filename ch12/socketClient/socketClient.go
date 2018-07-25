package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a socket file")
		os.Exit(1)
	}

	socketFile := os.Args[1]

	c, err := net.Dial("unix", socketFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()

	go readSocket(c)
	for {
		_, err := c.Write([]byte("Hello server!"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		time.Sleep(1 * time.Second)
	}
}

func readSocket(r io.Reader) {

	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("-> ", string(buf[0:n]))

	}
}
