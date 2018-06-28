package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func addIntegers(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum += input
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

func sendIntegers(c chan int, upper int) {
	for {
		c <- rand.Intn(upper)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <n>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	n, _ := strconv.Atoi(os.Args[1])
	c := make(chan int)
	go sendIntegers(c, n)
	go addIntegers(c)
	time.Sleep(2 * time.Second)
}
