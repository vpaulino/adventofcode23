package main

import "testing"

func TestCountDigits(t *testing.T) {

	cases := []struct {
		input    string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"treb7uchet78", 78},
	}

	for _, testCase := range cases {
		t.Run(testCase.input, func(t *testing.T) {
			result, error := ParseAndConvertToInt(testCase.input)
			if error != nil || result != testCase.expected {
				t.Errorf("Expected %d digits, but got %d ", testCase.expected, result)
			}

		})
	}
}