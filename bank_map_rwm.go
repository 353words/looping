package main

import "fmt"

type Account struct {
	Name    string
	Type    string
	Balance int
}

func main() {
	bank := map[string]Account{
		"donald":  {"donald", "regular", 123},
		"scrooge": {"scrooge", "vip", 1_000_001},
	}

	for l, a := range bank {
		if a.Type == "vip" {
			a.Balance += 1_000
			bank[l] = a
		}
	}
	fmt.Println(bank)
}
