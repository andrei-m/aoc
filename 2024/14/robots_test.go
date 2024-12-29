package main

import (
	"testing"

	"github.com/andrei-m/aoc/advent"
	"github.com/stretchr/testify/assert"
)

func Test_robot_advance(t *testing.T) {
	t.Run("example scenario", func(t *testing.T) {
		origHeight := height
		origWidth := width
		defer func() {
			height = origHeight
			width = origWidth
		}()

		width = 11
		height = 7
		r := robot{
			pos: advent.Point{X: 2, Y: 4},
			vel: advent.Point{X: 2, Y: -3},
		}

		r.advance()
		assert.Equal(t, advent.Point{X: 4, Y: 1}, r.pos)
		r.advance()
		assert.Equal(t, advent.Point{X: 6, Y: 5}, r.pos)
		r.advance()
		assert.Equal(t, advent.Point{X: 8, Y: 2}, r.pos)
		r.advance()
		assert.Equal(t, advent.Point{X: 10, Y: 6}, r.pos)
		r.advance()
		assert.Equal(t, advent.Point{X: 1, Y: 3}, r.pos)
	})

}
