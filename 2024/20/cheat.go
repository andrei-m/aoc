package main

import (
	"bufio"
	"fmt"
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
	part1(maze, startPos, endPos)
}

func part1(maze [][]object, startPos advent.Point, endPos advent.Point) {
	/*
		part 1:
			disabling colisions still costs movement: 1 picosecond per d-pad move, including through what are normally walls.
			1. Build a node adjacency graph for each unblocked D-pad direction starting from S
			2. Run Djikstra's algorithm to find the shortest path from S to E. Each edge cost is 1. The number of picoseconds is the baseline
			3. For each wall inside the perimeter on rows & columns with indices 0 and 140: Replace the wall with a path and re-run steps 1 and 2. If the shortest path is 100 or fewer than the baseline, count the wall as a cheat.
	*/

	graph := getPathGraph(maze, startPos)
	shortestPath := advent.GetShortestPath(graph, startPos)
	distance := advent.TraverseShortestPath(shortestPath, endPos)
	log.Printf("baseline distance: %d", distance)

	cheatCount := 0
	// omit border walls from consideration for cheat paths
	for y := 1; y < len(maze)-1; y++ {
		for x := 1; x < len(maze[0])-1; x++ {
			if maze[y][x] != wall {
				continue
			}
			maze[y][x] = path
			//TODO: it would be more efficient to modify only the adjacents of (x, y) in original graph, then reset that state at the end (rather than rebuilding the whole graph for each iteration)
			// removeWall(x, y) -> adjacents can navigate to (x, y); (x, y) can navigate to adjacents of (x y)
			// placeWall(x, y) -> adjacents of (x, y) cannot navigate to (x, y); remove (x, y) as a key from the adjacency graph
			cheatGraph := getPathGraph(maze, startPos)
			cheatPath := advent.GetShortestPath(cheatGraph, startPos)
			cheatDist := advent.TraverseShortestPath(cheatPath, endPos)
			if cheatDist+100 <= distance {
				cheatCount++
			}

			log.Printf("wall at %s: distance %d, cheat count: %d", advent.Point{X: x, Y: y}, cheatDist, cheatCount)
			maze[y][x] = wall
		}
	}

	fmt.Printf("part 1: %d\n", cheatCount)
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
