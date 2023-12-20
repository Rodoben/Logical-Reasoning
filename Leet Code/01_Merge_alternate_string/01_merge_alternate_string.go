package main

import "fmt"

// word1 --> abcd
// word2 --> efghjk
// output --> aebfcgdhjk

func main() {
	fmt.Println(mergeAlternateString("abcd", "zxcvaa"))
}

func mergeAlternateString(word1, word2 string) string {
	var result string
	l := len(word1)
	if len(word2) > l {
		l = len(word2)
	}
	for i := 0; i < l; i++ {
		if i < len(word1) {
			result += string(word1[i])
		}

		if i < len(word2) {
			result += string(word2[i])
		}
	}
	return result
}
