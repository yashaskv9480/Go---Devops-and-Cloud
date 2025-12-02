package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func Perform_http_call() {
	// URL := os.Getenv("URL")
	URL := "https://dummyjson.com/products"

	// url, err := url.ParseRequestURI(URL)

	if url, err := url.ParseRequestURI(URL); err != nil {
		fmt.Print("%s is not found or wrong", url)
		os.Exit(1)
	}

	response, _ := http.Get(URL)             //stream response
	ioOutput, _ := io.ReadAll(response.Body) // Real all to buffer
	StringOutput := string(ioOutput)         // Convert to String
	fmt.Print(StringOutput)
}
