package main

import (
	"fmt"
)

func main() {
	str1 := "armrrronald"
	str2 := "ronald"

	// b := strings.Contains(str1, str2)
	// fmt.Println(b)
	// b1 := strings.Index(str1, str2)
	// fmt.Println(b1)

	fmt.Println(findSubstringIndex(str1, str2))
}

func findSubstringIndex(str1, str2 string) int {
	for i := 0; i <= len(str1); i++ {
		match := true
		for j := 0; j < len(str2); j++ {
			fmt.Println("here", str1[i+j], str2[j])
			if str1[i+j] != str2[j] {
				match = false
				//fmt.Println("when", i, j)
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
