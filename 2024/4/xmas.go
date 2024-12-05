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

	solve1(puzzle)
	solve2(puzzle)
}

func solve1(puzzle [][]string) {
	matches := [][]int{}
	count := 0
	for y, row := range puzzle {
		for col := range row {
			if puzzle[y][col] == "X" {
				newMatches, newCount := masCount(puzzle, part1, y, col)
				if newCount > 0 {
					matches = append(matches, []int{y, col})
					matches = append(matches, newMatches...)
					count += newCount
				}
			}
		}
	}

	debug := make([][]string, len(puzzle))
	for i := range puzzle {
		debug[i] = make([]string, len(puzzle[i]))

		for j := range puzzle[i] {
			debug[i][j] = "."
		}
	}
	for _, match := range matches {
		debug[match[0]][match[1]] = puzzle[match[0]][match[1]]
	}

	for _, row := range debug {
		fmt.Println(strings.Join(row, ""))
	}

	fmt.Printf("count: %d", count)
}

func solve2(puzzle [][]string) {
	matches := [][]int{}
	count := 0
	for y, row := range puzzle {
		for col := range row {
			if puzzle[y][col] == "A" {
				newMatches, newCount := masCount(puzzle, part2, y, col)
				if newCount > 0 {
					matches = append(matches, []int{y, col})
					matches = append(matches, newMatches...)
					count += newCount
				}
			}
		}
	}

	debug := make([][]string, len(puzzle))
	for i := range puzzle {
		debug[i] = make([]string, len(puzzle[i]))

		for j := range puzzle[i] {
			debug[i][j] = "."
		}
	}
	for _, match := range matches {
		debug[match[0]][match[1]] = puzzle[match[0]][match[1]]
	}

	for _, row := range debug {
		fmt.Println(strings.Join(row, ""))
	}

	fmt.Printf("count: %d\n", count)
}

type vector struct {
	rowOffset int
	colOffset int
	char      string
}

// all vectors in a 'MAS' match need to hit for the match to be considered
type masMatch []vector

// each of the opportunities countes as one 'MAS' count
var (
	part1 = []masMatch{
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

	part2 = []masMatch{
		// MAS / MAS
		{
			{-1, -1, "M"},
			{1, 1, "S"},
			{-1, 1, "M"},
			{1, -1, "S"},
		},
		// SAM / MAS
		{
			{-1, -1, "S"},
			{1, 1, "M"},
			{-1, 1, "M"},
			{1, -1, "S"},
		},
		// SAM / SAM
		{
			{-1, -1, "S"},
			{1, 1, "M"},
			{-1, 1, "S"},
			{1, -1, "M"},
		},
		// MAS / SAM
		{
			{-1, -1, "M"},
			{1, 1, "S"},
			{-1, 1, "S"},
			{1, -1, "M"},
		},
	}
)

func masCount(puzzle [][]string, opportunities []masMatch, row int, col int) ([][]int, int) {
	matches := [][]int{}
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
			for _, v := range opt {
				matches = append(matches, []int{row + v.rowOffset, col + v.colOffset})
			}
		}
	}

	return matches, count
}
