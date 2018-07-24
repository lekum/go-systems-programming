package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide a socket file.")
		os.Exit(100)
	}
	socketFile := os.Args[1]
	l, err := net.Listen("unix", socketFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		go echoServer(fd)
	}
}

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 1024)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		fmt.Printf("->: %v\n", string(data))
		_, err = c.Write(data)
		if err != nil {
			fmt.Println(err)
		}
	}

}
