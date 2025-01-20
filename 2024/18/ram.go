package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	corruptions := mustParseInput(os.Stdin)
	log.Printf("%d points", len(corruptions))

	/*
		part 1:
		- plot the first N (1024) points on a 71x71 grid (0-70 are valid coordinate values). Each point represents corrupted memory.
		- Model a graph relation starting at (0,0) to span all possible traversals fo of the map. A traversal is possible in D-pad directions unless the adjacent square is corrupted memory.
		- Use Djikstra's algorithm to find a shortest path
	*/
	part1(corruptions)
}

func part1(corruptions []advent.Point) {
	xyOverflow := 71
	adjacencts = advent.AdjacentsFn(xyOverflow, xyOverflow)

	start := advent.Point{X: 0, Y: 0}
	pathGraph := getPathGraph(corruptions[:1024], start)
	shortestPath := getShortestPath(pathGraph, start)
	steps := traverseShortestPath(shortestPath, advent.Point{X: xyOverflow - 1, Y: xyOverflow - 1})
	fmt.Printf("part 1: %d\n", steps)
}

type edge struct {
	cost         int
	previousNode advent.Point
}

const inf = math.MaxInt

func getShortestPath(pathGraph map[advent.Point][]advent.Point, start advent.Point) map[advent.Point]edge {
	distances := map[advent.Point]edge{}
	toVisit, visited := map[advent.Point]struct{}{}, map[advent.Point]struct{}{}

	for _, adjacents := range pathGraph {
		for _, adj := range adjacents {
			distances[adj] = edge{cost: inf}
			toVisit[adj] = struct{}{}
		}
	}
	distances[start] = edge{cost: 0}
	toVisit[start] = struct{}{}

	for {
		if len(toVisit) == 0 {
			break
		}

		visitNextCost := inf
		var visitNext advent.Point
		for n := range toVisit {
			if distances[n].cost < visitNextCost {
				visitNextCost = distances[n].cost
				visitNext = n
			}
		}

		adjacents := pathGraph[visitNext]
		for _, adj := range adjacents {
			_, previouslyVisited := visited[adj]
			if previouslyVisited {
				continue
			}

			cost := 1 + visitNextCost
			if cost < distances[adj].cost {
				distances[adj] = edge{cost: cost, previousNode: visitNext}
			}
		}

		delete(toVisit, visitNext)
		visited[visitNext] = struct{}{}
	}

	return distances
}

func traverseShortestPath(shortestPath map[advent.Point]edge, end advent.Point) int {
	pos := end
	steps := 0
	for {
		if pos.X == 0 && pos.Y == 0 {
			return steps
		}
		pos = shortestPath[pos].previousNode
		steps++
	}
}

func getPathGraph(corruptions []advent.Point, start advent.Point) map[advent.Point][]advent.Point {
	corruptionsLookup := map[advent.Point]struct{}{}
	for _, c := range corruptions {
		corruptionsLookup[c] = struct{}{}
	}
	log.Printf("%d corruption obstacles present", len(corruptionsLookup))

	g := map[advent.Point][]advent.Point{}
	toVisit := []advent.Point{start}
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
		for _, adj := range adjacencts(current) {
			_, corrupt := corruptionsLookup[adj]
			if corrupt {
				continue
			}
			g[current] = append(g[current], adj)
			toVisit = append(toVisit, adj)
		}
	}
	return g
}

var adjacencts func(pos advent.Point) []advent.Point

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
