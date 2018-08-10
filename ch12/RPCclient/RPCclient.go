package main

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/lekum/go-systems-programming/ch12/sharedRPC"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide a host:port string")
		os.Exit(1)
	}

	addr := os.Args[1]

	c, err := rpc.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args := sharedRPC.MyInts{7, 18, true, false}
	var reply int

	err = c.Call("MyInterface.Add", args, &reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Reply (Add): %d\n", reply)

	err = c.Call("MyInterface.Substract", args, &reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Reply (Substract): %d\n", reply)

}
