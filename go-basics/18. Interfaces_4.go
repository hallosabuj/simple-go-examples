package main

import "fmt"

func main() {
	var rw ReaderWriter = MyReaderWriter{} // This is not possible with pointer receiver
	rw.Read()
	rw.Write()
	var rw1 ReaderWriter = &MyReaderWriter{} // This is possible with non pointer receiver
	rw1.Read()
	rw1.Write()
	// Conversion to MyReaderWriter
	mrw := rw.(MyReaderWriter)
	mrw.Read()
	mrw.Write()
}

type Reader interface {
	Read()
}
type Writer interface {
	Write()
}

// Composing smaller interfaces to make one larger interface
type ReaderWriter interface {
	Reader
	Writer
}

type MyReaderWriter struct{}

func (rw MyReaderWriter) Read() {
	fmt.Println("My read function")
}
func (rw MyReaderWriter) Write() {
	fmt.Println("My write function")
}
