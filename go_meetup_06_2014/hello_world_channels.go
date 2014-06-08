package main

import (
	"fmt"
	"time"
)

func hello(c chan string) {
	for{
		c <- fmt.Sprintf("Hello world.\n")
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go hello(c)
	for i := 0; i < 5; i++ {
		fmt.Printf("From a channel: %v", <-c)
	}
}
