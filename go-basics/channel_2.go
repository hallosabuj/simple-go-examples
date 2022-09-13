package main

import (
	"fmt"
	"strconv"
)

func myFunc(ch chan string) {
	for i := 0; i < 4; i++ {
		ch <- "Value : " + strconv.Itoa(i)
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	go myFunc(ch)
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("No more value to return")
			break
		}
		fmt.Println(value)
	}
}
