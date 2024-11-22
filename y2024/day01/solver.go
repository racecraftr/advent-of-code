package main

import (
	"fmt"
	"io"
	"net/http"
)

type Result[T any] struct {
	data T
	err error
}

func (r *Result[T]) Unwrap() (T, error) {
	return r.data, r.err
}

func part1() { 
	
}

func part2() {
	
}

func getBody(url string) (string, error) {
	
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	
	body, err := io.ReadAll(resp.Body)
	if (err != nil) {
		return "", err
	}
	return string(body), nil
}

func thisIsCool () {
	c := make(chan struct{url string; result Result[string]})
	urls := []string{
		"https://www.google.com",
		"https://thisisnotawebsitedotcom.com/",
		"this should be an error",
	}
	for _, url := range urls {
		go func(s string) {
			body, err := getBody(s)
			c <- struct{url string; result Result[string]}{s, Result[string]{body, err}}
		}(url)
	}
	
	for range urls {
		urlResult := <- c
		
		url, result := urlResult.url, urlResult.result
		body, err := result.Unwrap()
		
		if (err == nil) {
			fmt.Printf("%v has content:\n%v\n", url, body)
		} else {
			fmt.Printf("Could not get content for %v\n", url)
		}
		println()
	}
}

func main() {
	
}
