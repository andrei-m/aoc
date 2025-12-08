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

func Test_remainderRepeatsSequence(t *testing.T) {
	assert.True(t, remainderRepeatsSequence("0", "0"))
	assert.True(t, remainderRepeatsSequence("0", "00"))
	assert.True(t, remainderRepeatsSequence("12", "12"))
	assert.True(t, remainderRepeatsSequence("12", "1212"))
	assert.False(t, remainderRepeatsSequence("1", "2"))
	assert.False(t, remainderRepeatsSequence("11", "111"))
	assert.True(t, remainderRepeatsSequence("1234", "123412341234"))
	assert.False(t, remainderRepeatsSequence("1234", "1234123412345"))
	assert.False(t, remainderRepeatsSequence("", ""))
}

func Test_isValid(t *testing.T) {
	assert.True(t, isValid(123))
	assert.False(t, isValid(11))
	assert.False(t, isValid(22))
	assert.False(t, isValid(999))
	assert.False(t, isValid(1010))
	assert.True(t, isValid(1011))
	assert.False(t, isValid(1188511885))
}

func Test_getInvalidIDsP2(t *testing.T) {
	assert.Equal(t, []int{11, 22}, getInvalidIDsP2(idRange{11, 22}))
	assert.Equal(t, []int{2121212121}, getInvalidIDsP2(idRange{2121212118, 2121212124}))
}
