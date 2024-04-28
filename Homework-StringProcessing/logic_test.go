package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    string
	}{
		{"1. empty", "", ""},
		{"2. one letter", "A", "A"},
		{"3. two letters", "AC", "AC"},
		{"4. two letters - BA", "BA", ""},
		{"5. two letters - AB", "AB", ""},
		{"6. several letters - combinations", "BAAB", ""},
		{"7. several letters", "CBACD", "C"},
		{"8. several letters", "CABABD", ""},
		{"9. several letters", "ACBDACBD", "ACBDACBD"},
		{"10. several letters", "ABA", "A"},
		{"11. several letters", "ABACD", "A"},
	}

	for _, tc := range cases {
		t.Run(tc.description,
			func(t *testing.T) {
				require.EqualValues(t,
					tc.expected,
					Parse(tc.input),
				)
			},
		)
	}
}
