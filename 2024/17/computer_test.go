package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_debug(t *testing.T) {
	f, err := os.Open("/path/to/input.txt")
	require.NoError(t, err)
	defer f.Close()

	c := mustParseInput(f)
	part1(c)
}
