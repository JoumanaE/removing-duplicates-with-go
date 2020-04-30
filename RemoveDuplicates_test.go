package main

import (
	"testing"
)

type tt struct {
	bcc      []string
	email    string
	expected int
}

func TestBcc(t *testing.T) {
	tabletests := []tt{
		{bcc: []string{"test@test.com"},
			email:    "testing@boring.com",
			expected: 1,
		},
	}

	for _, tabletest := range tabletests {
		output := RemoveDuplicates(tabletest.email, tabletest.bcc)
		if len(output) != tabletest.expected {
			t.Errorf("I got %v, and I expected %v ", len(output), tabletest.expected)

		}

	}

}
