package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_find(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		result := find(1, []int{})
		assert.Equal(t, -1, result)
	})

	t.Run("element not found", func(t *testing.T) {
		result := find(2, []int{1, 17, 99, 103, 104, 107})
		assert.Equal(t, -1, result)
	})

	t.Run("element found", func(t *testing.T) {
		result := find(104, []int{1, 17, 99, 103, 104, 107})
		assert.Equal(t, 4, result)
	})

	t.Run("element found longer list", func(t *testing.T) {
		result := find(104, []int{1, 17, 18, 21, 99, 100, 103, 104, 107, 109, 112, 113, 120})
		assert.Equal(t, 7, result)
	})

	t.Run("element exists multiple times", func(t *testing.T) {
		result := find(17, []int{1, 17, 17, 99, 103, 104, 107})
		assert.True(t, result == 1 || result == 2)
	})
}

func Test_countOccurrences(t *testing.T) {
	t.Run("one occurrence", func(t *testing.T) {
		count := countOccurrences(0, []int{0, 1, 2, 3})
		assert.Equal(t, 1, count)
	})

	t.Run("later occurrence", func(t *testing.T) {
		count := countOccurrences(1, []int{0, 1, 1, 2, 3})
		assert.Equal(t, 2, count)
	})

	t.Run("earlier occurrence", func(t *testing.T) {
		count := countOccurrences(6, []int{0, 1, 1, 2, 3, 3, 3})
		assert.Equal(t, 3, count)

	})

	t.Run("both earlier and later occurrences", func(t *testing.T) {
		count := countOccurrences(4, []int{0, 1, 1, 2, 2, 2, 2, 3, 3, 3})
		assert.Equal(t, 4, count)
	})
}
