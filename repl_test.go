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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "hello , world",
			expected: []string{"hello", ",", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, v := range cases {
		actual := cleanInput(v.input)
		if len(actual) != len(v.expected) {
			t.Errorf("Wrong length\ninput: %s\nexpected length:%d\nactual length:%d", v.input, len(v.expected), len(actual))
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			if expectedWord := v.expected[i]; expectedWord != word {
				t.Errorf("Wrong word\ninput: %s\nexpected:%s\nactual:%s", v.input, expectedWord, word)
				t.Fail()
			}
		}
	}

}
