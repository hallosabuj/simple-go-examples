package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop() // CPU Profiling
	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop() // Memory Profiling
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop() // Memory Profiling
	fmt.Println("Hello")
	f, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	words := 0
	inword := false
	b := bufio.NewReader(f)
	for {
		r, err := readbyte(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Could not read file")
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Println("text.txt : ", words)
}
