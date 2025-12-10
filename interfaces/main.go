package main

import "io"

func main() {
	s := StringContents{
		contents: "hii",
	}

	io.ReadAll(s)
	PrintandScan(s)
}

func PrintandScan(p PrintScanner) {
	p.Print()
	p.Scan()
}
