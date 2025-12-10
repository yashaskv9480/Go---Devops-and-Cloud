package main

import "fmt"

type RequestError struct {
	HTTPCode int
	Body     string
	Err      string
}

func (r RequestError) Error() string {
	return fmt.Sprintf("Error: %s | Status: %d | Body: %s", r.Err, r.HTTPCode, r.Body)
}
