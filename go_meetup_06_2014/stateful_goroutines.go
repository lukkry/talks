package main

import (
	"fmt"
	"math/rand"
	"time"
)

// types start OMIT
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

// types end OMIT

func main() {
	// channels start OMIT
	reads := make(chan *readOp)
	writes := make(chan *writeOp)
	// channels end OMIT
	quit := make(chan bool)

	go func() {
		state := make(map[int]int)
		readCount := 0
		writeCount := 0

		for {
			select {
			case read := <-reads:
				readCount += 1
				read.resp <- state[read.key]
			case write := <-writes:
				writeCount += 1
				state[write.key] = write.val
				write.resp <- true
			case <-quit:
				fmt.Printf("Reads: %d\n", readCount)
				fmt.Printf("Writes: %d", writeCount)
				return
			}
		}
	}()
	// state OMIT

	// read start OMIT
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
			}
		}()
	}
	// read end OMIT

	// write start OMIT
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
			}
		}()
	}
	// write end OMIT

	fmt.Println("Sleeping for a second...")
	time.Sleep(time.Second)
	quit <- true
}
