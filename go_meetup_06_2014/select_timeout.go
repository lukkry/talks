package main

import (
	"fmt"
	"time"
)

func hello(c chan string) {
	for{
		c <- fmt.Sprintf("Hello world\n")
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	timeout := time.After(2 * time.Second)
	go hello(c)
	for {
		select {
		case s := <-c:
			fmt.Printf("From a channel: %v", s)
		case <-timeout:
			fmt.Println("Timeout.")
			return
		}
	}
}
