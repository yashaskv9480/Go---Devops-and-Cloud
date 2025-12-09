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

type GithubResponse struct {
	Name  string         `json:"full_name"`
	Owner map[string]any `json:"owner"`
}

func Perform_http_call() {
	// URL := os.Getenv("URL")
	URL := "https://catfact.ninja /fact"
	GITHUBURL := "https://api.github.com/repos/golang/go"

	if url, err := url.ParseRequestURI(URL); err != nil {
		fmt.Print("%s is not found or wrong", url)
		os.Exit(1)
	}

	// StringOutput := string(ioOutput)         // Convert to String

	// log.Println("String is ", StringOutput)

	githubReponse, err := getResponse(GITHUBURL)
	urlResponse, err := getResponse(URL)

	var jsonOutput CatResponse
	var githubOutput GithubResponse

	err = json.Unmarshal(urlResponse, &jsonOutput) //Convert to Json Object
	err = json.Unmarshal(githubReponse, &githubOutput)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GITHUB Json Result is %s and %f", githubOutput.Name, githubOutput.Owner["id"])
	fmt.Printf("JSON result is %v", jsonOutput)

}

func getResponse(URL string) ([]byte, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	ioOutput, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return ioOutput, nil
}
