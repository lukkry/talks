package main

import (
	"fmt"
	"time"
	"math/rand"
)

func download(filenames []string) []string{
	results := make([]string, len(filenames))
	for i := 0; i < len(filenames); i++ {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		results = append(results, filenames[i])
	}
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := download([]string{"1.mp3", "2.mp3"})
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
