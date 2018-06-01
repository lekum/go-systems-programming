package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	if len(os.Args) == 1 {
		myString = "You did not give an argument!"
	} else {
		myString = os.Args[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
