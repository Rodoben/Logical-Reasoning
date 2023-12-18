package main

import (
	"reflect"
	"testing"
)

func Test_Fibonaaci(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedOutput []int
	}{
		{name: "valid1", input: 5, expectedOutput: []int{0, 1, 1, 2, 3}},
		{name: "valid2", input: 6, expectedOutput: []int{0, 1, 1, 2, 3, 5}},
	}

	for _, test := range tests {
		fib := fibonacci(test.input)
		if !reflect.DeepEqual(fib, test.expectedOutput) {
			t.Errorf("%s: expected %v:, got %v:", test.name, test.expectedOutput, fib)
		}
	}
}
