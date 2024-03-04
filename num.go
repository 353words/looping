package main

import (
	"fmt"
	"testing"
)

func main() {
	for i := range 3 {
		fmt.Println(i)
	}
}

type Event struct{}

func NewEvent(login, kind, uri string) Event {
	return Event{}
}

func (Event) Validate() error { return nil }

func BenchmarkEvent_Validae(b *testing.B) {
	evt := NewEvent("elliot", "read", "/etc/passwd")
	for range b.N {
		err := evt.Validate()
		if err != nil {
			b.Fatal(err)
		}
	}
}
