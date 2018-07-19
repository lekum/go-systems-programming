package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Telephone struct {
	Mobile bool
	Number string
}

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]

	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

}
