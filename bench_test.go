package main

import "testing"

func Benchmarkname(b *testing.B) {
	var data = `{"login": "joe", "path": "/users", "method": "POST"}`
	for range b.N {
		err := Decode(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Decode(data string) error {
	return nil
}
