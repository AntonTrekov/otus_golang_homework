package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var (
		result      strings.Builder
		buffer      strings.Builder
		resultError error
	)

	for _, charRune := range input {
		if buffer.Len() == 0 && unicode.IsLetter(charRune) {
			buffer.WriteRune(charRune)

			continue
		}

		if buffer.Len() == 1 && unicode.IsLetter(charRune) {
			result.WriteString(buffer.String())
			buffer.Reset()
			buffer.WriteRune(charRune)

			continue
		}

		if buffer.Len() == 1 && unicode.IsDigit(charRune) {
			count, err := strconv.Atoi(string([]rune{charRune}))

			if err != nil {
				return "", err
			}

			result.WriteString(
				strings.Repeat(buffer.String(), count),
			)
			buffer.Reset()

			continue
		}

		resultError = ErrInvalidString
		result.Reset()
	}

	if resultError == nil {
		result.WriteString(buffer.String())
	}

	return result.String(), resultError
}
