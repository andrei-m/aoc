package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func Test_debug(t *testing.T) {
	f, err := os.Open("/path/to/input.txt")
	require.NoError(t, err)
	defer f.Close()

	c := mustParseInput(f)
	part1(&c)
}
*/

func Test_part1(t *testing.T) {
	t.Run("scenario 1", func(t *testing.T) {
		c := computer{state: state{regC: 9}, program: []int{2, 6}}
		part1(&c)
		assert.Equal(t, 1, c.regB)
	})

	t.Run("scenario 2", func(t *testing.T) {
		c := computer{state: state{regA: 10}, program: []int{5, 0, 5, 1, 5, 4}}
		part1(&c)
		assert.Equal(t, []int{0, 1, 2}, c.out)
	})

	t.Run("scenario 3", func(t *testing.T) {
		c := computer{state: state{regA: 2024}, program: []int{0, 1, 5, 4, 3, 0}}
		part1(&c)
		assert.Equal(t, []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}, c.out)
		assert.Equal(t, 0, c.regA)
	})

	t.Run("scenario 4", func(t *testing.T) {
		c := computer{state: state{regB: 29}, program: []int{1, 7}}
		part1(&c)
		assert.Equal(t, 26, c.regB)
	})

	t.Run("scenario 5", func(t *testing.T) {
		c := computer{state: state{regB: 2024, regC: 43690}, program: []int{4, 0}}
		part1(&c)
		assert.Equal(t, 44354, c.regB)
	})
}
