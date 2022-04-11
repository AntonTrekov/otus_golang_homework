package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	t.Run("test add", func(t *testing.T) {
		builder := Builder{}
		builder.Add([]string{"two", "one"})

		require.Equal(t, []string{"one", "two"}, builder.GetTopWords(10))

		builder.Add([]string{"three", "four"})

		require.Equal(t, []string{"four", "one", "three", "two"}, builder.GetTopWords(10))
	})

	t.Run("test add after reset", func(t *testing.T) {
		builder := Builder{}
		builder.Add([]string{"two", "one"}).Reset()

		require.Equal(t, []string{}, builder.GetTopWords(10))

		builder.Add([]string{"three", "four"})

		require.Equal(t, []string{"four", "three"}, builder.GetTopWords(10))
	})
}

func TestGetTopWords(t *testing.T) {
	t.Run("test get top words", func(t *testing.T) {
		builder := Builder{}

		require.Equal(t, []string{}, builder.GetTopWords(10))
		require.Equal(t, []string{}, builder.GetTopWords(0))
		require.Equal(t, []string{}, builder.GetTopWords(-1))

		builder.Add([]string{"two", "one"})

		require.Equal(t, []string{"one", "two"}, builder.GetTopWords(10))
		require.Equal(t, []string{}, builder.GetTopWords(0))
		require.Equal(t, []string{}, builder.GetTopWords(-1))
		require.Equal(t, []string{"one"}, builder.GetTopWords(1))
	})
}

func TestReset(t *testing.T) {
	t.Run("test reset", func(t *testing.T) {
		builder := Builder{}
		builder.Add([]string{"two", "one"}).Reset()

		require.Equal(t, []string{}, builder.GetTopWords(10))
	})
}
