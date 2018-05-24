package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("* Got", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handleSignal(sig)
			case syscall.SIGTERM:
				handleSignal(sig)
			case syscall.SIGHUP:
				fmt.Println("Got:", sig)
				os.Exit(-1)

			}
		}
	}()
	for {
		fmt.Printf(".")
		time.Sleep(10 * time.Second)
	}
}
