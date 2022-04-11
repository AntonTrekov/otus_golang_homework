package hw03frequencyanalysis

import (
	"sort"
)

type Builder struct {
	buf    []entry
	sorted bool
}

func (b *Builder) Add(items []string) *Builder {
	b.addEntries(items)
	b.sorted = false

	return b
}

func (b *Builder) GetTopWords(count int) []string {
	if count < 0 {
		count = 0
	}
	if len(b.buf) < count {
		count = len(b.buf)
	}
	if !b.sorted {
		b.sort()
	}

	result := make([]string, count)

	for i := 0; i < count; i++ {
		result[i] = b.buf[i].Word
	}

	return result
}

func (b *Builder) Reset() *Builder {
	b.buf = nil

	return b
}

func (b *Builder) addEntries(items []string) *Builder {
	entriesMap := make(map[string]int)

	if b.buf == nil {
		b.buf = make([]entry, 0)
	}

	for _, item := range items {
		if count, ok := entriesMap[item]; ok {
			entriesMap[item] = count + 1
		} else {
			entriesMap[item] = 1
		}
	}

	for word, count := range entriesMap {
		b.buf = append(b.buf, entry{word, count})
	}

	return b
}

func (b *Builder) sort() {
	sort.Slice(
		b.buf,
		func(i, j int) bool {
			if b.buf[i].Count == b.buf[j].Count {
				return b.buf[i].Word < b.buf[j].Word
			}

			return b.buf[i].Count > b.buf[j].Count
		},
	)

	b.sorted = true
}

type entry struct {
	Word  string
	Count int
}
