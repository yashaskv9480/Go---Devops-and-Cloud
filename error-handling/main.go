package main

import (
	"fmt"
	"io"
	"net/http"
)

func callAPI(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", RequestError{
			Err: "Network error: " + err.Error(), // Uses err.Error function
		}
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyStr := string(bodyBytes)

	if resp.StatusCode != 200 {
		return "", RequestError{
			HTTPCode: resp.StatusCode,
			Body:     bodyStr,
			Err:      "Non-200 response received ",
		}
	}

	return bodyStr, nil
}

func main() {
	result, err := callAPI("https://httpbin.org/status/500")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", result)
}
