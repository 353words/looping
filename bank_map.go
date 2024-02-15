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

	for l := range bank {
		fmt.Println(l)
	}

	for l, a := range bank {
		fmt.Printf("%s -> %+v\n", l, a)
	}

	for _, a := range bank {
		fmt.Printf("%+v\n", a)
	}

	for _, a := range bank {
		if a.Type == "vip" {
			a.Balance += 1_000
		}
	}
	fmt.Println(bank)

	for l, a := range bank {
		a.Balance += 1000
		bank[l] = a
	}
	fmt.Println(bank)
}
