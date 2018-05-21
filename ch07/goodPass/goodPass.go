package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"os"
)

const MAX int = 90
const MIN int = 0

const seedSize int = 10

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	minusL := flag.Int("l", 10, "Length of the generated password")
	flag.Parse()

	passwordLength := *minusL
	f, _ := os.Open("/dev/random")
	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	rand.Seed(seed)
	f.Close()
	fmt.Println("Seed:", seed)

	startChar := "!"
	var i int
	for i = 0; i < passwordLength; i++ {
		anInt := int(random(MIN, MAX))
		newChar := string(startChar[0] + byte(anInt))
		if newChar == " " {
			i = i - 1
			continue
		}
		fmt.Print(newChar)
	}
	fmt.Println()
}
