package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxJoltage(t *testing.T) {
	assert.Equal(t, 98, maxJoltage([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}))
	assert.Equal(t, 89, maxJoltage([]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}))
	assert.Equal(t, 78, maxJoltage([]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}))
	assert.Equal(t, 92, maxJoltage([]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}))
}
