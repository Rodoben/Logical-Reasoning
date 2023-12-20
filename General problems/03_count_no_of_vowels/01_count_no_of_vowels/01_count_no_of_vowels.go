package main

import (
	"fmt"
	"strings"
)

// []string{a,a,e,r,t,y,u,y,t,a}

func main() {

	vowelsUnclean := []string{"a", "a", "e", "r"}

	// warning never declare a map like this, nil map
	//var vowelsCleaned map[string]int

	// green flag of declaring a map
	//var vowelsCleaned = map[string]int{}
	var vowelsCleaned = make(map[string]int)

	//vowelsCleaned := map[string]int{}

	for _, v := range vowelsUnclean {
		if isVowel(v) {
			if _, ok := vowelsCleaned[v]; !ok {
				vowelsCleaned[v] = 1
			} else {
				vowelsCleaned[v]++
			}
		}
	}
	fmt.Println(vowelsCleaned)
	for i, v := range vowelsCleaned {
		fmt.Println(i, v)
	}
}

func isVowel(s string) bool {
	s = strings.ToLower(s)

	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		return true
	}
	return false
}
