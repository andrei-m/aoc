package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mustParseScenario(t *testing.T) {
	lines := []string{
		"Button A: X+1, Y+2",
		"Button B: X+3, Y+4",
		"Prize: X=55, Y=66",
	}
	scenario := mustParseScenario(lines)
	assert.Equal(t, 1, scenario.aDelta.X)
	assert.Equal(t, 2, scenario.aDelta.Y)
	assert.Equal(t, 3, scenario.bDelta.X)
	assert.Equal(t, 4, scenario.bDelta.Y)
	assert.Equal(t, 55, scenario.prizeLoc.X)
	assert.Equal(t, 66, scenario.prizeLoc.Y)
}
