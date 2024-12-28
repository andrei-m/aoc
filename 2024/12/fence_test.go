package main

import (
	"testing"

	"github.com/andrei-m/aoc/advent"
	"github.com/stretchr/testify/assert"
)

func Test_scoreRegionPart2(t *testing.T) {
	t.Run("edge region", func(t *testing.T) {
		/*
			HHHH
			 HHH
			HH
			HH
			 H

			 area: 12
			 expected perimeter: 12
		*/
		adjacentDir = adjacentDirFn(144, 144)
		r := region{
			plant: "H",
			locs: []advent.Point{
				{X: 42, Y: 0},
				{X: 43, Y: 0},
				{X: 45, Y: 0},
				{X: 44, Y: 1},
				{X: 42, Y: 3},
				{X: 44, Y: 0},
				{X: 45, Y: 1},
				{X: 43, Y: 1},
				{X: 43, Y: 2},
				{X: 42, Y: 2},
				{X: 43, Y: 3},
				{X: 43, Y: 4},
			},
		}
		assert.Equal(t, 144, scoreRegionPart2(r))
	})

}
