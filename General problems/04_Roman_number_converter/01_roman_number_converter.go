package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConvertRomanToNumber("IX"))

}

func ConvertRomanToNumber(n string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
	}

	sum := 0
	prevValue := 0
	for _, v := range n {
		if i, ok := romans[string(v)]; ok {
			sum += i
			if prevValue < i {
				sum = sum - (2 * prevValue)
			}
			prevValue = i
		}
	}
	return sum
}
