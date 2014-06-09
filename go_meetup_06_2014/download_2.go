package main

import (
	"fmt"
	"time"
	"math/rand"
)

func download(filename string) string{
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("LOL# %v\n", filename)
	return filename
}

func fanIn(filenames []string) []string {
	results := make([]string, len(filenames))
	c := make(chan string)

	//go func(){ c <- download(filenames[0]) }()
	//go func(){ c <- download(filenames[1]) }()

	//for i := 0; i < len(filenames); i++ {
		//fmt.Printf("LOL %v\n", i)
		//go func(){ c <- download(filenames[i]) }()
	//}

	for _, v := range filenames {
		fmt.Printf("LOL %v\n", v)
		go func(){ c <- download(v) }()
	}

	for i := 0; i < len(filenames); i++ {
		result := <-c
		fmt.Printf("LOL@ %v\n", result)
		results = append(results, result)
	}

	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := fanIn([]string{"1.mp3", "2.mp3"})
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
