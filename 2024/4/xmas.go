package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	puzzle := [][]string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		puzzle = append(puzzle, strings.Split(scanner.Text(), ""))
	}

	count := 0
	for row := range puzzle {
		for col := range row {
			if puzzle[row][col] == "X" {
				count += masCount(puzzle, row, col)
			}
		}
	}

	fmt.Printf("part 1: %d", count)
}

type vector struct {
	rowOffset int
	colOffset int
	char      string
}

// all vectors in a 'MAS' match need to hit for the match to be considered
type masMatch []vector

// each of the opportunities countes as one 'MAS' count
var opportunities = []masMatch{
	{
		// right
		{1, 0, "M"},
		{2, 0, "A"},
		{3, 0, "S"},
	},
	{
		// left
		{-1, 0, "M"},
		{-2, 0, "A"},
		{-3, 0, "S"},
	},
	{
		// down
		{0, 1, "M"},
		{0, 2, "A"},
		{0, 3, "S"},
	},
	{
		// up
		{0, -1, "M"},
		{0, -2, "A"},
		{0, -3, "S"},
	},
	{
		// down/right
		{1, 1, "M"},
		{2, 2, "A"},
		{3, 3, "S"},
	},
	{
		// up/left
		{-1, -1, "M"},
		{-2, -2, "A"},
		{-3, -3, "S"},
	},
	{
		// down/left
		{1, -1, "M"},
		{2, -2, "A"},
		{3, -3, "S"},
	},
	{
		// up/right
		{-1, 1, "M"},
		{-2, 2, "A"},
		{-3, 3, "S"},
	},
}

func masCount(puzzle [][]string, row int, col int) int {
	count := 0
	for _, opt := range opportunities {
		eligible := true

		for _, v := range opt {
			y := row + v.rowOffset
			x := col + v.colOffset

			if y < 0 || y > len(puzzle)-1 {
				eligible = false
				// row out of bounds
				break
			}
			if x < 0 || x > len(puzzle[y])-1 {
				eligible = false
				// column out of bounds
				break
			}

			if puzzle[y][x] != v.char {
				eligible = false
				// no character match
				break
			}
		}

		if eligible {
			count++
		}
	}

	return count
}
