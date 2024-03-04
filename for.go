package main

import (
	"fmt"
)

func main() {
	n := 3
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	for i := range n {
		fmt.Printf("%d ", i)
	}

	fmt.Println()
	fmt.Println(IsPalindrome("()()")) // false
	fmt.Println(IsPalindrome("())(")) // true

	fmt.Println()
	// Sum of even Fibonacci numbers up to 4_000_000
	a, b, total := 1, 2, 0
	for a <= 4_000_000 {
		if a%2 == 0 {
			total += a
		}
		a, b = b, a+b
	}
	fmt.Println("total:", total)

	fmt.Println()
	handler(Provider{})

	fmt.Println()
	for i := range 3 {
		fmt.Print(i)
	}
}

func handler(p Provider) {
	for {
		msg := p.Next()
		if msg == nil {
			break
		}
		// TODO: Handle message
	}
}

func (p *Provider) Next() *Message {
	if p.n == 3 {
		return nil
	}

	p.n++
	return &Message{}
}

type Message struct{}
type Provider struct {
	n int
}

// IsPalindrome returns true of `s` is a palindrome.
func IsPalindrome(s string) bool {
	rs := []rune(s) // convert to runes
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if rs[i] != rs[j] {
			return false
		}
	}

	return true
}
