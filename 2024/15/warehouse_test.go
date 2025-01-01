package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func Test_debugPart2(t *testing.T) {

	f, err := os.Open("/path/to/input")
	require.NoError(t, err)
	defer f.Close()

	m, dirs := mustParseInput(f)
	part2(m, dirs)
}
*/

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
