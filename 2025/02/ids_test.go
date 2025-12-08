package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getInvalidIDs(t *testing.T) {
	assert.Equal(t, []int{11, 22}, getInvalidIDs(idRange{11, 22}))
	assert.Equal(t, []int{99}, getInvalidIDs(idRange{95, 115}))
	assert.Equal(t, []int{1010}, getInvalidIDs(idRange{998, 1012}))
	assert.Equal(t, []int{1188511885}, getInvalidIDs(idRange{1188511880, 1188511890}))
}
