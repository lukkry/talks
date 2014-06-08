package main

import (
	"fmt"
	"time"
)

func hello() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello world #%d\n", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go hello()
	time.Sleep(2 * time.Second)
}
