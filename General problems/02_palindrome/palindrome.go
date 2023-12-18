package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(palindrome("naman"))
	fmt.Println(palindrome("A man a plan a canal Panama"))
	fmt.Println(palindrome("A man a plan a canal Panam"))
	fmt.Println(palindrome("Ronald"))
}

// naman
func palindrome(s string) bool {
	s = strings.ToLower(strings.Join(strings.Fields(s), ""))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
