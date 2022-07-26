package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(Data []byte) (int, error) {
	n, err := fmt.Println(string(Data))
	return n, err
}

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World"))
}
