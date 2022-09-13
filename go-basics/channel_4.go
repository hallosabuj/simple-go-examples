package main

import "fmt"

func main() {
	ch := make(chan string, 5)
	ch <- "Message 1"
	ch <- "Message 2"
	ch <- "Message 3"
	ch <- "Message 4"

	fmt.Println(len(ch), " items in the channel")
}
