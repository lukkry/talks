package main

import (
	"fmt"
	"time"
	"math/rand"
)

func hello(name string) chan string{
	c := make(chan string)
	go func() {
		for i := 0; ; i++{
			c <- fmt.Sprintf("%v: Hello %v.\n", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	joe := hello("joe")
	ann := hello("ann")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v", <-joe)
		fmt.Printf("%v", <-ann)
	}
}
