package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

func readSocket(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, _ := r.Read(buf[:])
		fmt.Print("Read: ", string(buf[0:n]))
	}
}

func main() {
	c, err := net.Dial("unix", "/tmp/aSocket.sock")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening socket: %s\n", err)
		os.Exit(1)
	}
	defer c.Close()

	go readSocket(c)
	n := 0

	for {
		message := []byte("Hi there: " + strconv.Itoa(n) + "\n")
		_, err = c.Write(message)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to socket: %s\n", err)
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
		n++
	}
}
