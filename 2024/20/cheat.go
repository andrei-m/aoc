package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	maze := mustParseInput(os.Stdin)
	yOverflow, xOverflow := len(maze), len(maze[0])
	fmt.Printf("%dx%d grid", xOverflow, yOverflow)

	/*
		part 1:
			disabling colisions still costs movement: 1 picosecond per d-pad move, including through what are normally walls.
			1. Build a node adjacency graph for each unblocked D-pad direction starting from S
			2. Run Djikstra's algorithm to find the shortest path from S to E. Each edge cost is 1. The number of picoseconds is the baseline
			3. For each wall inside the perimeter on rows & columns with indices 0 and 140: Replace the wall with a path and re-run steps 1 and 2. If the shortest path is 100 or fewer than the baseline, count the wall as a cheat.
	*/
}

type object int

const (
	path object = iota
	wall
	start
	end
)

func mustParseInput(r io.Reader) [][]object {
	scanner := bufio.NewScanner(r)
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
			case "E":
				row[i] = end
			default:
				log.Fatalf("invalid char: %s", line[i])
			}
		}
		rows = append(rows, row)
	}

	return rows
}
