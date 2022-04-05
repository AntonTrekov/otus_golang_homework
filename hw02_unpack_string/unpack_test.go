package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "AbC2", expected: "AbCC"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `чебур2ек`, expected: `чебуррек`},
		{input: `💩4💩`, expected: `💩💩💩💩💩`},
		{input: `]*#2@`, expected: `]*##@`},
		{input: `牡マ2キ`, expected: `牡ママキ`},
		{input: `ل2لأعيان`, expected: `لللأعيان`},
		{input: `b.3`, expected: `b...`},
		{input: `"check"`, expected: `"check"`},
		{input: `t 3st`, expected: `t   st`},
		{input: `\t\r`, expected: `tr`},
		{input: `	2a	`, expected: `		a	`},
		{input: `
2`, expected: `

`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `\\qwe`, expected: `\qwe`},
		{input: `\\2qwe`, expected: `\\qwe`},
		{input: `\\\32qwe`, expected: `\33qwe`},
		{input: `\1qwe`, expected: `1qwe`},
		{input: `\12qwe`, expected: `11qwe`},
		{input: `\1\\2qwe`, expected: `1\\qwe`},
		{input: `\1`, expected: `1`},
		{input: `\\`, expected: `\`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []struct {
		input string
	}{
		{input: `3abc`},
		{input: `45`},
		{input: `aaa10b`},
	}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			_, err := Unpack(tc.input)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
