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
		err               error
		isEscapingEnabled = false
	)

	for _, charRune := range input {
		if buffer.Len() == 0 {
			err = processFistSymbolInBlock(charRune, &isEscapingEnabled, &buffer)

			if err == nil {
				continue
			}
		}

		if buffer.Len() == 1 {
			err = processSecondSymbolInBlock(charRune, &isEscapingEnabled, &buffer, &result)

			if err == nil {
				continue
			}
		}

		if buffer.Len() > 1 {
			err = ErrInvalidString
		}

		if err != nil {
			result.Reset()
			break
		}
	}

	if err == nil {
		result.WriteString(buffer.String())
	}

	return result.String(), err
}

func processFistSymbolInBlock(charRune rune, isEscapingEnabled *bool, buffer *strings.Builder) error {
	if charRune == backslashRune && !*isEscapingEnabled {
		*isEscapingEnabled = true

		return nil
	}

	if unicode.IsLetter(charRune) || isSymbolEscaped(charRune, *isEscapingEnabled) {
		buffer.WriteRune(charRune)
		*isEscapingEnabled = false

		return nil
	}

	return ErrInvalidString
}

func processSecondSymbolInBlock(
	charRune rune,
	isEscapingEnabled *bool,
	buffer *strings.Builder,
	result *strings.Builder,
) error {
	if charRune == backslashRune {
		result.WriteString(buffer.String())
		buffer.Reset()
		*isEscapingEnabled = true

		return nil
	}

	if unicode.IsLetter(charRune) || isSymbolEscaped(charRune, *isEscapingEnabled) {
		result.WriteString(buffer.String())
		buffer.Reset()
		*isEscapingEnabled = false
		buffer.WriteRune(charRune)

		return nil
	}

	if buffer.Len() == 1 && unicode.IsDigit(charRune) {
		count, err := strconv.Atoi(string([]rune{charRune}))
		if err != nil {
			result.Reset()

			return err
		}

		result.WriteString(
			strings.Repeat(buffer.String(), count),
		)
		buffer.Reset()
		*isEscapingEnabled = false

		return nil
	}

	return ErrInvalidString
}

func isSymbolEscaped(charRune rune, isEscapingEnabled bool) bool {
	return isEscapingEnabled && (unicode.IsDigit(charRune) || charRune == backslashRune)
}
