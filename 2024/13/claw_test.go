package main

import (
	"testing"

	"github.com/andrei-m/aoc/advent"
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

func Test_getSolution(t *testing.T) {
	t.Run("scenario 1", func(t *testing.T) {
		// scenario: {{14 98} {34 36} {296 1062}}; solutions: [{9 5}]: best score: 32
		scen := scenario{
			aDelta:   advent.Point{X: 14, Y: 98},
			bDelta:   advent.Point{X: 34, Y: 36},
			prizeLoc: advent.Point{X: 296, Y: 1062},
		}
		sol := getSolution(scen, 100)
		if assert.NotNil(t, sol) {
			assert.Equal(t, solution{a: 9, b: 5}, *sol)
		}
	})

	t.Run("scenario 2", func(t *testing.T) {
		// {{20 42} {43 12} {3325 4172}}; no solution
		scen := scenario{
			aDelta:   advent.Point{X: 20, Y: 42},
			bDelta:   advent.Point{X: 43, Y: 12},
			prizeLoc: advent.Point{X: 3325, Y: 4172},
		}
		assert.Nil(t, getSolution(scen, 100))
	})
}
