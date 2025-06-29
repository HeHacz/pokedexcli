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
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " justGo  ",
			expected: []string{"justgo"},
		},
		{
			input:    " GO GO golang rangers  ",
			expected: []string{"go", "go", "golang", "rangers"},
		},
		{
			input:    " Charmander bulbasaur SQUIRTLE ",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words, got %d words.", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word '%s', got '%s'.", expectedWord, word)
			}
		}
	}
}
