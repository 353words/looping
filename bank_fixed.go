package main

import "fmt"

type Account struct {
	Name    string
	Type    string
	Balance int
}

func main() {
	bank := []Account{
		{"donald", "regular", 123},
		{"scrooge", "vip", 1_000_001},
	}

	for _, a := range bank {
		if a.Type == "vip" {
			a.Balance += 1_000
		}
	}
	fmt.Println(bank)
}
