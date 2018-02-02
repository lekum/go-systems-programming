package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("Mihalis", "Mihalis Tsoukalos")
	fmt.Println(match)
	match, _ = regexp.MatchString("Tsoukalos", "Mihalis tsoukalos")
	fmt.Println(match)
	fmt.Println()

	parse, err := regexp.Compile("[Mm]ihalis")
	if err != nil {
		log.Fatalf("Error compiling RE: %s\n", err)
	} else {
		fmt.Println(parse.MatchString("Mihalis Tsoukalos"))
		fmt.Println(parse.MatchString("mihalis Tsoukalos"))
		fmt.Println(parse.MatchString("M ihalis Tsoukalos"))
		fmt.Println(parse.ReplaceAllString("mihalis Mihalis", "MIHALIS"))
	}
}
