package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0

	for _, arg := range os.Args[1:] {
		temp, err := strconv.Atoi(arg)
		if err != nil {
			continue
		}
		sum += temp
	}
	fmt.Println("Sum:", sum)
}
