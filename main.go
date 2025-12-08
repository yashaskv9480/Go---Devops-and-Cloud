package main

import "io"

func main() {
	// Perform_http_call()
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
