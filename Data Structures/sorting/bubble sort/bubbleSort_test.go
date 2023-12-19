package main

import (
	"reflect"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "valid",
			input:          []int{3, 1, 2},
			expectedOutput: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		xi := BubbleSort(test.input)
		if !reflect.DeepEqual(xi, test.expectedOutput) {
			t.Errorf("%s: expected %v, got %v", test.name, test.expectedOutput, xi)
		}
	}

}
