package main

import "testing"

// Test the code from day06/main.go
func TestDay06(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, test := range tests {
		result := getStartMarker([]byte(test.input))
		if result != test.expected {
			t.Errorf("getStartMarker(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
