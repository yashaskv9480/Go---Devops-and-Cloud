package main

import (
	"fmt"
	"io"
)

type PrintScanner interface {
	Print()
	Scan()
}

type StringContents struct { //Define a type
	contents string
}

func (s StringContents) Read(p []byte) (n int, err error) { // Add the interface methods
	fmt.Printf(s.contents)
	return 0, io.EOF
}

func (s StringContents) Print() {
	fmt.Print("Using Print")
}

func (s StringContents) Scan() {
	fmt.Print("Using Print")
}
