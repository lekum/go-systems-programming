package main

import (
	"fmt"
	"time"
)

func namedFunction() {
	time.Sleep(10000 * time.Microsecond)
	fmt.Println("Printing from namedFunction")
}

func main() {
	fmt.Println("Chapter 09 - Goroutines.")
	go namedFunction()
	go func() {
		fmt.Println("An anonymous function")
	}()
	time.Sleep(10000 * time.Microsecond)
	fmt.Println("Exiting...")
}
