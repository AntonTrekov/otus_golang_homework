package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	ErrInvalidString = errors.New("invalid string")
	backslash        = `\`
	backslashRune, _ = utf8.DecodeRuneInString(backslash)
)

func Unpack(input string) (string, error) {
	var (
		result            strings.Builder
		buffer            strings.Builder
		resultError       error
		isEscapingEnabled = false
	)

	for _, charRune := range input {
		if buffer.Len() == 0 && charRune == backslashRune && !isEscapingEnabled {
			isEscapingEnabled = true

			continue
		}

		isCurrentSymbolEscaped := isEscapingEnabled && (unicode.IsDigit(charRune) || charRune == backslashRune)

		if buffer.Len() == 0 && (unicode.IsLetter(charRune) || isCurrentSymbolEscaped) {
			buffer.WriteRune(charRune)
			isEscapingEnabled = false

			continue
		}

		if buffer.Len() == 1 && charRune == backslashRune {
			result.WriteString(buffer.String())
			buffer.Reset()
			isEscapingEnabled = true

			continue
		}

		if buffer.Len() == 1 && (unicode.IsLetter(charRune) || isCurrentSymbolEscaped) {
			result.WriteString(buffer.String())
			buffer.Reset()
			isEscapingEnabled = false
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
			isEscapingEnabled = false

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
