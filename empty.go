package main

import (
	"errors"
	"fmt"
	"net/http"
)

func Retry(fn func() (*http.Response, error), count int) (*http.Response, error) {
	var errs error

	for range count {
		resp, err := fn()
		if err == nil {
			return resp, nil
		}

		errs = errors.Join(errs, err)
	}

	return nil, errs
}

func main() {
	fn := func() (*http.Response, error) {
		return http.Get("https://go.devxxx")
	}
	fmt.Println(Retry(fn, 3))
}
