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

	for i, a := range bank {
		if a.Type == "vip" {
			a.Balance += 1_000
			bank[i] = a
		}
	}
	fmt.Println(bank)
}
