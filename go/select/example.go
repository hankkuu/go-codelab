package main

import (
	"fmt"
	"time"
)

func main() {

	in := make(chan string, 10)
	out := make(chan int)

	go Worker(in, out)
	go Print(out)

	for i := 0; i < 100000; i++ {
		in <- "https://example.com"
	}
	time.Sleep(20 * time.Second)

}

func Worker(in chan string, out chan int) {
	for val := range in {
		go Request(val, out)
	}

}

func Print(in chan int) {

	for val := range in {
		fmt.Println(val)
	}

}

func Request(val string, out chan int) {
	fmt.Println(val)
	time.Sleep(1 * time.Second) // request 보냄
	out <- 200
}
