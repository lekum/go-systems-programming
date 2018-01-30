package main

import (
	"fmt"
	"reflect"
)

type message struct {
	X     int
	Y     int
	Label string
}

func main() {

	p1 := message{23, 12, "A Message"}
	s1 := reflect.ValueOf(&p1).Elem()

	typeOfT := s1.Type()
	fmt.Println("P1=", p1)
	for i := 0; i < s1.NumField(); i++ {
		f := s1.Field(i)
		fmt.Printf("%d: %s ", i, typeOfT.Field(i).Name)
		fmt.Printf("%s = %v\n", f.Type(), f.Interface())
	}
}
