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
	maxlength := len(word1)
	if len(word2) > maxlength {
		maxlength = len(word2)
	}
	for i := 0; i < maxlength; i++ {
		if i < len(word1) {
			result += string(word1[i])
		}

		if i < len(word2) {
			result += string(word2[i])
		}
	}
	return result
}
