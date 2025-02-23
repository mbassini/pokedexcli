package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HeLlo  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello  world",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("\"Lengths don't match: '%v' vs '%v'\n", actual, c.expected)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%v) == %v, expected %v\n", c.input, actual[i], c.expected[i])
			}
		}
	}
}
