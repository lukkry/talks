package main

import (
	"fmt"
	"time"
	"math/rand"
)

func hello(name string) chan string{
	c := make(chan string)
	go func() {
		for {
			c <- fmt.Sprintf("%v: Hello.\n", name)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	joe := hello("joe")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v", <-joe)
	}
}
