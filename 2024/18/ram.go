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
		- Model a graph relation starting at (0,0) to span all possible traversals of of the map. A traversal is possible in D-pad directions unless the adjacent square is corrupted memory.
		- Use Djikstra's algorithm to find a shortest path
	*/
	part1(corruptions)

	/*
		part 2:
		- Starting at 1024 build the the traversal adjacency graph
		- attempt DFS or BFS from (0,0) to (70, 70)
		- if found, apppend the next corruption and try again. If not found, return the most recently added corruption
	*/
	part2(corruptions)
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

func part2(corruptions []advent.Point) {
	xyOverflow := 71
	adjacencts = advent.AdjacentsFn(xyOverflow, xyOverflow)

	start := advent.Point{X: 0, Y: 0}
	end := advent.Point{X: 70, Y: 70}
	corruptionCount := 1024
	for i := corruptionCount; i <= len(corruptions); i++ {
		pathGraph := getPathGraph(corruptions[:i], start)
		if !reachable(pathGraph, start, end) {
			fmt.Printf("part 2: %s (%d corruptions)\n", corruptions[i-1], i)
			return
		}
	}

	fmt.Printf("exit still reachable after %d corruptions processed", corruptionCount)
}

func reachable(pathGraph map[advent.Point][]advent.Point, start advent.Point, end advent.Point) bool {
	toVisit := map[advent.Point]struct{}{
		start: {},
	}
	visited := map[advent.Point]struct{}{}

	for {
		if len(toVisit) == 0 {
			return false
		}
		var currentNode advent.Point
		for k := range toVisit {
			currentNode = k
			break
		}
		delete(toVisit, currentNode)

		nextNodes := pathGraph[currentNode]
		for _, n := range nextNodes {
			if n == end {
				return true
			}

			_, previouslyVisited := visited[n]
			if !previouslyVisited {
				toVisit[n] = struct{}{}
			}
		}
		visited[currentNode] = struct{}{}
	}
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
