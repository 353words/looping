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

	for k := range bank {
		if bank[k].Type == "vip" {
			bank[k].Balance += 1_000
		}
	}
	fmt.Println(bank)

}
