package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RotateClockwise(t *testing.T) {
	assert.Equal(t, Down, RotateClockwise(Right))
	assert.Equal(t, Left, RotateClockwise(Down))
	assert.Equal(t, Up, RotateClockwise(Left))
	assert.Equal(t, Right, RotateClockwise(Up))
}

func Test_RotateCounterClockwise(t *testing.T) {
	assert.Equal(t, Up, RotateCounterClockwise(Right))
	assert.Equal(t, Right, RotateCounterClockwise(Down))
	assert.Equal(t, Down, RotateCounterClockwise(Left))
	assert.Equal(t, Left, RotateCounterClockwise(Up))
}
