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
func palindrome(str string) bool {
	str = strings.ToLower(strings.Join(strings.Fields(str), ""))
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
