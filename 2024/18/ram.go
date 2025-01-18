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
	points := mustParseInput(os.Stdin)
	log.Printf("%d points", len(points))

	/*
		part 1:
		- plot the first N (1024) points on a 71x71 grid (0-70 are valid coordinate values). Each point represents corrupted memory.
		- Model a graph relation starting at (0,0) to span all possible traversals fo of the map. A traversal is possible in D-pad directions unless the adjacent square is corrupted memory.
		- Use Djikstra's algorithm to find a shortest path
	*/
}

func mustParseInput(r io.Reader) []advent.Point {
	points := []advent.Point{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ",")
		if len(lineParts) != 2 {
			log.Fatalf("invalid line: %s", line)
		}
		p := advent.Point{X: advent.MustParseInt(lineParts[0]), Y: advent.MustParseInt(lineParts[1])}
		points = append(points, p)
	}
	return points
}
