package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

//https://mholt.github.io/json-to-go/
type CatResponse struct { //For json parse
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func Perform_http_call() {
	// URL := os.Getenv("URL")
	URL := "https://catfact.ninja/fact"

	// url, err := url.ParseRequestURI(URL)

	if url, err := url.ParseRequestURI(URL); err != nil {
		fmt.Print("%s is not found or wrong", url)
		os.Exit(1)
	}

	response, err := http.Get(URL) //stream response

	if err != nil || response.StatusCode != 200 {
		var logInfo any

		if err != nil {
			logInfo = err
		} else {
			logInfo = response.Body
		}
		log.Fatal(logInfo) //exits
	}
	fmt.Print("hii")
	ioOutput, _ := io.ReadAll(response.Body) // Read all to buffer
	StringOutput := string(ioOutput)         // Convert to String

	log.Println("String is ", StringOutput)

	var jsonOutput CatResponse
	err = json.Unmarshal(ioOutput, &jsonOutput) //Convert to Json Object

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Json Result is ", jsonOutput)
	defer response.Body.Close()
}
