package main

import (
	"reflect"
	"testing"
)

func Test_FibonaacciConcurrent(t *testing.T) {

	tests := []struct {
		name           string
		input          int
		expectedOutput []int
	}{
		{name: "valid", input: 6, expectedOutput: []int{0, 1, 1, 2, 3, 5}},
	}

	for _, test := range tests {
		xi := make([]int, 0)
		for n := range fibonacciConcurrent(test.input) {
			xi = append(xi, n)
		}

		if !reflect.DeepEqual(xi, test.expectedOutput) {
			t.Errorf("%s: expected %v:, got %v:", test.name, test.expectedOutput, xi)
		}
	}
}
