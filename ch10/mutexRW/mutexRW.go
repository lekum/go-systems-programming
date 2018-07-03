package main

import (
	"fmt"
	"sync"
	"time"
)

var Password = secret{counter: 1, password: "myPassword"}

type secret struct {
	sync.RWMutex
	counter  int
	password string
}

func (c *secret) Change(pass string) {
	c.Lock()
	fmt.Println("LChange")
	time.Sleep(20 * time.Second)
	c.counter += 1
	c.password = pass
	c.Unlock()
}

func (c *secret) Show() string {
	fmt.Println("LShow")
	time.Sleep(time.Second)
	c.RLock()
	defer c.RUnlock()
	return c.password
}

func (c *secret) Counts() int {
	c.RLock()
	defer c.RUnlock()
	return c.counter
}

func main() {
	fmt.Println("Pass:", Password.Show())
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Go Pass:", Password.Show())
		}()
	}
	go func() {
		Password.Change("123456")
	}()
	fmt.Println("Pass:", Password.Show())
	time.Sleep(time.Second)
	fmt.Println("Counter:", Password.Counts())
}
