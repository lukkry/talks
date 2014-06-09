package main

import (
	"fmt"
	"time"
)

func hello(c chan string) {
	for i := 0; ; i++{
		c <- fmt.Sprintf("Hello %d.\n", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go hello(c)
	for i := 0; i < 5; i++ {
		fmt.Printf("Channel: %v", <-c)
	}
	fmt.Printf("Ok, I'm done")
}
