package main

import "testing"

func Test_Palindrome(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput bool
	}{
		{name: "valid", input: "naman", expectedOutput: true},
	}

	for _, test := range tests {
		v := palindrome(test.input)

		if v != test.expectedOutput {
			t.Errorf("%s: expected output: %v, but got: %v", test.name, test.expectedOutput, v)
		}
	}
}
