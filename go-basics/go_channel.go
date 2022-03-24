package main

import (
	"fmt"
	"time"
)

var (
	i int32
)

func worker(ch chan struct{}) {
	for {
		select {
		case <-ch:
			i++
		}
	}
}

var counter1 int = 0
var counter2 int = 0

func func1(ch chan struct{}) {
	for {
		counter1++
		ch <- struct{}{}
	}
}
func func2(ch chan struct{}) {
	for {
		counter2++
		ch <- struct{}{}
	}
}

func main() {
	ch := make(chan struct{})
	go func1(ch)
	go func2(ch)
	go worker(ch)
	time.Sleep(10 * time.Second)

	fmt.Println("Final value of i is :", i, counter1, counter2)
}
