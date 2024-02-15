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

}

// IsPalindrome returns true of `s` is a palindrome.
func IsPalindrome(s string) bool {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if rs[i] != rs[j] {
			return false
		}
	}

	return true
}
