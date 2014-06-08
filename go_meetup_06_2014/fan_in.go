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
			c <- fmt.Sprintf("%v: Hello world.\n", name)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 chan string) chan string {
	c := make(chan string)
	go func(){ for { c <- <-input1 }}()
	go func(){ for { c <- <-input2 }}()
	return c
}

func main() {
	c := fanIn(hello("joe"), hello("ann"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%v", <-c)
	}
}
