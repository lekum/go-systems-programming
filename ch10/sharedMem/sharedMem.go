package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func SetValue(newValue int) {
	writeValue <- newValue
}

func ReadValue() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	go monitor()

	var wg sync.WaitGroup

	for r := 0; r < 20; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			SetValue(rand.Intn(100))
		}()
	}

	wg.Wait()
	fmt.Printf("\nLast value: %d\n", ReadValue())
}
