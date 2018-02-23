package main

import (
	"fmt"
	"github.com/lekum/go-systems-programming/ch04/aSimplePackage"
)

func main() {
	temp := aSimplePackage.Add(5, 10)
	fmt.Println(temp)
	fmt.Println(aSimplePackage.Pi)
	aSimplePackage.Version()
	fmt.Println(aSimplePackage.version)
}
