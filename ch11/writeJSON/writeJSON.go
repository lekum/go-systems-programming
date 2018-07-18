package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

func saveToJSON(filename string, key interface{}) {
	out, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s\n", err)
		return
	}
	defer out.Close()

	encodeJSON := json.NewEncoder(out)
	err = encodeJSON.Encode(key)
	if err != nil {
		fmt.Printf("Error encoding JSON: %s\n", err)
		return
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage %s <filename>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	filename := os.Args[1]
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{Telephone{Mobile: true, Number: "1234-567"},
			Telephone{Mobile: true, Number: "1234-abcd"},
			Telephone{Mobile: false, Number: "abcc-567"},
		}}

	saveToJSON(filename, myRecord)
}
