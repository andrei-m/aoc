package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	maze, startPos := mustParseInput(os.Stdin)
	yOverflow, xOverflow := len(maze), len(maze[0])

	var endPos advent.Point
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == end {
				endPos = advent.Point{X: x, Y: y}
			}
		}
	}
	log.Printf("%dx%d grid; start: %s; end: %s", xOverflow, yOverflow, startPos, endPos)

	adjacents = advent.AdjacentsFn(xOverflow, yOverflow)
	graph := getPathGraph(maze, startPos)

	shortestPath := advent.GetShortestPath(graph, startPos)
	log.Printf("baseline distance: %d", advent.TraverseShortestPath(shortestPath, endPos))

	/*
		part 1:
			disabling colisions still costs movement: 1 picosecond per d-pad move, including through what are normally walls.
			1. Build a node adjacency graph for each unblocked D-pad direction starting from S
			2. Run Djikstra's algorithm to find the shortest path from S to E. Each edge cost is 1. The number of picoseconds is the baseline
			3. For each wall inside the perimeter on rows & columns with indices 0 and 140: Replace the wall with a path and re-run steps 1 and 2. If the shortest path is 100 or fewer than the baseline, count the wall as a cheat.
	*/
}

func getPathGraph(maze [][]object, startPos advent.Point) map[advent.Point][]advent.Point {
	g := map[advent.Point][]advent.Point{}

	toVisit := []advent.Point{startPos}

	for {
		if len(toVisit) == 0 {
			break
		}
		current := toVisit[0]
		toVisit = toVisit[1:]

		_, visited := g[current]
		if visited {
			continue
		}

		for _, adj := range adjacents(current) {
			if maze[adj.Y][adj.X] != wall {
				g[current] = append(g[current], adj)
				toVisit = append(toVisit, adj)
			}
		}
	}

	return g
}

var adjacents func(advent.Point) []advent.Point

type object int

const (
	path object = iota
	wall
	start
	end
)

func mustParseInput(r io.Reader) ([][]object, advent.Point) {
	scanner := bufio.NewScanner(r)
	var startPos advent.Point
	rows := [][]object{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		row := make([]object, len(line))
		for i := range line {
			switch line[i] {
			case ".":
				row[i] = path
			case "#":
				row[i] = wall
			case "S":
				row[i] = start
				startPos = advent.Point{X: i, Y: len(rows)}
			case "E":
				row[i] = end
			default:
				log.Fatalf("invalid char: %s", line[i])
			}
		}
		rows = append(rows, row)
	}

	return rows, startPos
}
