package main

import (
	"fmt"
	"time"
	"math/rand"
)

func hello() {
	for i := 0; ; i++ {
		fmt.Printf("Hello %d\n", i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func main() {
	go hello()
}
