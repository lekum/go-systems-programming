package main

import (
	"errors"
	"fmt"
	"log"
)

func division(x, y int) (int, error, error) {
	if y == 0 {
		return 0, nil, errors.New("Cannot divide by zero!")
	}
	if x%y != 0 {
		remainder := errors.New("There is a remainder!")
		return x / y, remainder, nil
	} else {
		return x / y, nil, nil
	}
}

func main() {
	r, err1, err2 := division(3, 0)
	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println("Result", r)
	}
}
