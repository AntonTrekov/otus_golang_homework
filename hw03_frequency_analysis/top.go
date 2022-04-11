package hw03frequencyanalysis

import (
	"strings"
)

func Top10(input string) []string {
	if input == "" {
		return make([]string, 0)
	}

	words := explode(input)
	builder := Builder{}

	return builder.Add(words).GetTopWords(10)
}

func explode(input string) []string {
	return strings.Fields(input)
}
