package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_score(t *testing.T) {
	m := [][]object{
		{wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall},
		{wall, wall, boxLeft, boxRight, empty, empty, empty, empty, empty, empty, empty, boxLeft, boxRight, empty, boxLeft, boxRight, boxLeft, boxRight, wall, wall},
		{wall, wall, boxLeft, boxRight, empty, empty, empty, empty, empty, empty, empty, empty, empty, empty, empty, boxLeft, boxRight, empty, wall, wall},
		{wall, wall, boxLeft, boxRight, empty, empty, empty, empty, empty, empty, empty, empty, boxLeft, boxRight, boxLeft, boxRight, boxLeft, boxRight, wall, wall},
		{wall, wall, boxLeft, boxRight, empty, empty, empty, empty, empty, empty, boxLeft, boxRight, empty, empty, empty, empty, boxLeft, boxRight, wall, wall},
		{wall, wall, empty, empty, wall, wall, empty, empty, empty, empty, empty, empty, boxLeft, boxRight, empty, empty, empty, empty, wall, wall},
		{wall, wall, empty, empty, boxLeft, boxRight, empty, empty, empty, empty, empty, empty, empty, empty, empty, empty, empty, empty, wall, wall},
		{wall, wall, empty, empty, robot, empty, empty, empty, empty, empty, empty, boxLeft, boxRight, empty, boxLeft, boxRight, boxLeft, boxRight, wall, wall},
		{wall, wall, empty, empty, empty, empty, empty, empty, boxLeft, boxRight, boxLeft, boxRight, empty, empty, boxLeft, boxRight, empty, empty, wall, wall},
		{wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall, wall},
	}
	assert.Equal(t, 9021, score(m))
}

func Test_part2(t *testing.T) {
	t.Run("edge case 1", func(t *testing.T) {
		in := `#######
#.....#
#.OO@.#
#.....#
#######

<<`
		r := strings.NewReader(in)
		m, dirs := mustParseInput(r)
		s := part2(m, dirs)
		assert.Equal(t, 406, s)
	})

	t.Run("edge case 2", func(t *testing.T) {
		in := `#######
#.....#
#.O#..#
#..O@.#
#.....#
#######

<v<<^`
		r := strings.NewReader(in)
		m, dirs := mustParseInput(r)
		s := part2(m, dirs)
		assert.Equal(t, 509, s)
	})

	t.Run("edge case 3", func(t *testing.T) {
		in := `#######
#.....#
#.#O..#
#..O@.#
#.....#
#######

<v<^`
		r := strings.NewReader(in)
		m, dirs := mustParseInput(r)
		s := part2(m, dirs)
		assert.Equal(t, 511, s)
	})

	t.Run("edge case 4", func(t *testing.T) {
		in := `######
#....#
#.O..#
#.OO@#
#.O..#
#....#
######

<vv<<^`
		r := strings.NewReader(in)
		m, dirs := mustParseInput(r)
		s := part2(m, dirs)
		assert.Equal(t, 816, s)
	})

	t.Run("edge case 5", func(t *testing.T) {
		in := `#######
#...#.#
#.....#
#.....#
#.....#
#.....#
#.OOO@#
#.OOO.#
#..O..#
#.....#
#.....#
#######

v<vv<<^^^^^`
		r := strings.NewReader(in)
		m, dirs := mustParseInput(r)
		s := part2(m, dirs)
		assert.Equal(t, 2339, s)
	})
}
