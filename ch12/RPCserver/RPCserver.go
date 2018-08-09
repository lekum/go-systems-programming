package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"

	"github.com/lekum/go-systems-programming/ch12/sharedRPC"
)

type MyInterface int

func (t *MyInterface) Add(arguments *sharedRPC.MyInts, reply *int) error {
	s1 := 1
	s2 := 1

	if arguments.S1 == true {
		s1 = -1
	}

	if arguments.S2 == true {
		s2 = -1
	}

	*reply = s1*int(arguments.A1) + s2*int(arguments.A2)

	return nil
}

func (t *MyInterface) Substract(arguments *sharedRPC.MyInts, reply *int) error {
	s1 := 1
	s2 := 1

	if arguments.S1 == true {
		s1 = -1
	}

	if arguments.S2 == true {
		s2 = -1
	}

	*reply = s1*int(arguments.A1) - s2*int(arguments.A2)

	return nil

}

func main() {

	port := ":1234"

	myInterface := new(MyInterface)
	rpc.Register(myInterface)

	t, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l, err := net.ListenTCP("tcp", t)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(c)
	}

}
