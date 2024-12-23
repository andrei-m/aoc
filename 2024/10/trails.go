package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	/*
		Parse input file into a graph (represented as a map of nodes to adjacent node references)
		each node is identified by an (x, y) coordinate and has a topgraphic value (0-9)
		iterate over each node; if it's a 0, perform DFS of adjacent nodes, with the topgraphic height increasing by 1 at each hop, until a height value 9 is found. If found, count the 9 as reachable from that trailhead.
		- Don't count the same node twice within the same trailhead.
	*/
	rows := mustParseIntRows()
	trailheads := getTrailheads(rows)

	var overallScore int
	for _, t := range trailheads {
		score := score(rows, t)
		overallScore += score
		log.Printf("%v: %d", t, score)
	}
	fmt.Printf("part 1: %d\n", overallScore)

	overallScore = 0
	for _, t := range trailheads {
		score := scoreSearchPart2(rows, t)
		overallScore += score
		log.Printf("%v: %d", t, score)
	}
	fmt.Printf("part 2: %d\n", overallScore)
}

func score(rows [][]int, trailhead advent.Point) int {
	return len(scoreSearch(rows, trailhead))
}

func scoreSearch(rows [][]int, pos advent.Point) map[advent.Point]struct{} {
	currentHeight := rows[pos.Y][pos.X]
	if currentHeight == 9 {
		return map[advent.Point]struct{}{
			pos: {},
		}
	}

	nines := map[advent.Point]struct{}{}
	if pos.Y > 0 && rows[pos.Y-1][pos.X] == currentHeight+1 {
		nines = mergeMaps(nines, scoreSearch(rows, advent.Point{X: pos.X, Y: pos.Y - 1}))
	}
	if pos.Y < len(rows)-1 && rows[pos.Y+1][pos.X] == currentHeight+1 {
		nines = mergeMaps(nines, scoreSearch(rows, advent.Point{X: pos.X, Y: pos.Y + 1}))
	}
	if pos.X > 0 && rows[pos.Y][pos.X-1] == currentHeight+1 {
		nines = mergeMaps(nines, scoreSearch(rows, advent.Point{X: pos.X - 1, Y: pos.Y}))
	}
	if pos.X < len(rows[0])-1 && rows[pos.Y][pos.X+1] == currentHeight+1 {
		nines = mergeMaps(nines, scoreSearch(rows, advent.Point{X: pos.X + 1, Y: pos.Y}))
	}

	return nines
}

func scoreSearchPart2(rows [][]int, pos advent.Point) int {
	currentHeight := rows[pos.Y][pos.X]
	if currentHeight == 9 {
		return 1
	}

	var up, down, left, right int
	if pos.Y > 0 && rows[pos.Y-1][pos.X] == currentHeight+1 {
		up = scoreSearchPart2(rows, advent.Point{X: pos.X, Y: pos.Y - 1})
	}
	if pos.Y < len(rows)-1 && rows[pos.Y+1][pos.X] == currentHeight+1 {
		down = scoreSearchPart2(rows, advent.Point{X: pos.X, Y: pos.Y + 1})
	}
	if pos.X > 0 && rows[pos.Y][pos.X-1] == currentHeight+1 {
		left = scoreSearchPart2(rows, advent.Point{X: pos.X - 1, Y: pos.Y})
	}
	if pos.X < len(rows[0])-1 && rows[pos.Y][pos.X+1] == currentHeight+1 {
		right = scoreSearchPart2(rows, advent.Point{X: pos.X + 1, Y: pos.Y})
	}

	return up + down + left + right
}

func mergeMaps(a, b map[advent.Point]struct{}) map[advent.Point]struct{} {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func getTrailheads(rows [][]int) []advent.Point {
	trailheads := []advent.Point{}
	for y := range rows {
		for x := range rows[y] {
			if rows[y][x] == 0 {
				trailheads = append(trailheads, advent.Point{X: x, Y: y})
			}
		}
	}
	return trailheads
}

func mustParseIntRows() [][]int {
	rows := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		row := make([]int, len(line))
		for x, chr := range line {
			height, err := strconv.Atoi(chr)
			if err != nil {
				log.Fatalf("not an int %s: %v", chr, err)
			}
			row[x] = height
		}
		rows = append(rows, row)
	}
	return rows
}
