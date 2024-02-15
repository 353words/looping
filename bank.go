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

	for i := range bank {
		fmt.Println(i)
	}

	for i, a := range bank {
		fmt.Printf("%d: %+v\n", i, a)
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

	for i := range bank {
		if bank[i].Type == "vip" {
			bank[i].Balance += 1_000
		}
	}
	fmt.Println(bank)

}
