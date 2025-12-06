package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	newPos, zeroClicks := rotate(0, 0)
	assert.Equal(t, 0, newPos)
	assert.Equal(t, 0, zeroClicks)

	newPos, zeroClicks = rotate(50, -1)
	assert.Equal(t, 49, newPos)
	assert.Equal(t, 0, zeroClicks)

	newPos, zeroClicks = rotate(50, 1)
	assert.Equal(t, 51, newPos)
	assert.Equal(t, 0, zeroClicks)

	newPos, zeroClicks = rotate(1, -1)
	assert.Equal(t, 0, newPos)
	assert.Equal(t, 1, zeroClicks)

	newPos, zeroClicks = rotate(1, -2)
	assert.Equal(t, 99, newPos)
	assert.Equal(t, 1, zeroClicks)

	newPos, zeroClicks = rotate(99, 1)
	assert.Equal(t, 0, newPos)
	assert.Equal(t, 1, zeroClicks)

	newPos, zeroClicks = rotate(99, 2)
	assert.Equal(t, 1, newPos)
	assert.Equal(t, 1, zeroClicks)

	t.Run("part2 example", func(t *testing.T) {
		newPos, zeroClicks = rotate(50, -68)
		assert.Equal(t, 82, newPos)
		assert.Equal(t, 1, zeroClicks)

		newPos, zeroClicks = rotate(82, -30)
		assert.Equal(t, 52, newPos)
		assert.Equal(t, 0, zeroClicks)

		newPos, zeroClicks = rotate(52, 48)
		assert.Equal(t, 0, newPos)
		assert.Equal(t, 1, zeroClicks)

		newPos, zeroClicks = rotate(0, -5)
		assert.Equal(t, 95, newPos)
		assert.Equal(t, 0, zeroClicks)

		newPos, zeroClicks = rotate(95, 60)
		assert.Equal(t, 55, newPos)
		assert.Equal(t, 1, zeroClicks)

		newPos, zeroClicks = rotate(55, -55)
		assert.Equal(t, 0, newPos)
		assert.Equal(t, 1, zeroClicks)

		newPos, zeroClicks = rotate(0, -1)
		assert.Equal(t, 99, newPos)
		assert.Equal(t, 0, zeroClicks)

		newPos, zeroClicks = rotate(99, -99)
		assert.Equal(t, 0, newPos)
		assert.Equal(t, 1, zeroClicks)

		newPos, zeroClicks = rotate(0, 14)
		assert.Equal(t, 14, newPos)
		assert.Equal(t, 0, zeroClicks)

		newPos, zeroClicks = rotate(14, -82)
		assert.Equal(t, 32, newPos)
		assert.Equal(t, 1, zeroClicks)

		//2400 too low
	})
}
