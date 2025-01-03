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
	m, start, end := mustParseInput(os.Stdin)
	adjacencts = advent.AdjacentsFn(len(m[0]), len(m))
	log.Printf("%dx%d maze; start: %v; end: %v", len(m[0]), len(m), start, end)

	/*
		Use advent.Direction for geographic directions (advent.Right=="East" etc)
		From the starting position & Direction, build a graph of traversable paths. Include the traversal direction in each edge.
		Use Djikstra's algorithm with the following distance function:
		- Traversal in the same direction == 1
		- Traversal in a 90 degree direction == 1001
	*/

	pathGraph := getPathGraph(m, start)

	path := getShortestPaths(pathGraph, start, advent.Right)

	printMap(m, path, start, end)
	printPath(path, end)
}

// Build a map of destination->traversal cost/previous path step for each destination reachable from 'start' using Djikstra's algorithm
func getShortestPaths(pathGraph map[advent.Point][]advent.Point, start advent.Point, initialDir advent.Direction) map[advent.Point]*traversal {
	debugEnabled := advent.DebugEnabled()
	distances := map[advent.Point]*traversal{}
	toVisit := map[advent.Point]struct{}{}
	visited := map[advent.Point]struct{}{}

	for node := range pathGraph {
		distances[node] = &traversal{cost: inf}
		toVisit[node] = struct{}{}
	}
	// initial distances from 'start' (direction cannot be determined from previous node)
	distances[start].cost = 0

	// Visit each element of 'toVisit' in priority order, according to lowest cost in 'distances'
	for {
		if len(toVisit) == 0 {
			// visited all nodes
			break
		}

		lowestCost := inf
		var lowestCostNode advent.Point
		for loc := range toVisit {
			if distances[loc].cost < lowestCost {
				lowestCost = distances[loc].cost
				lowestCostNode = loc
			}
		}
		if debugEnabled {
			log.Printf("lowest cost node: %v", lowestCostNode)
		}

		lcnTraversal := distances[lowestCostNode]
		for i := range pathGraph[lowestCostNode] {
			adj := pathGraph[lowestCostNode][i]
			_, previouslyVisited := visited[adj]
			if previouslyVisited {
				continue
			}

			var edgeCost int
			var newDir advent.Direction
			if lcnTraversal.previousNode != nil {
				edgeCost, newDir = cost(lowestCostNode, adj, *lcnTraversal.dir)
			} else {
				// starting node
				edgeCost, newDir = cost(start, adj, initialDir)
			}

			if lcnTraversal.cost+edgeCost < distances[adj].cost {
				previousCost := distances[adj].cost
				distances[adj].cost = lcnTraversal.cost + edgeCost
				distances[adj].previousNode = &lowestCostNode
				distances[adj].dir = &newDir
				if debugEnabled {
					log.Printf("found lower cost %d (from %d) to get to %v (from %v; new direction: %v)", distances[adj].cost, previousCost, adj, lowestCostNode, newDir)
				}
			}
		}

		delete(toVisit, lowestCostNode)
		visited[lowestCostNode] = struct{}{}
	}

	return distances
}

// Calculate the cost and new direction of a potential edge traversal
func cost(previousNode advent.Point, nextNode advent.Point, previousDirection advent.Direction) (int, advent.Direction) {
	dX := nextNode.X - previousNode.X
	dY := nextNode.Y - previousNode.Y
	if advent.Abs(dX)+advent.Abs(dY) != 1 {
		log.Fatalf("%v and %v are not adjacent", previousNode, nextNode)
	}

	var newDirection advent.Direction
	if dX == -1 {
		newDirection = advent.Left
	} else if dX == 1 {
		newDirection = advent.Right
	} else if dY == -1 {
		newDirection = advent.Up
	} else if dY == 1 {
		newDirection = advent.Down
	} else {
		log.Fatalf("unexpected dX: %d or dY: %d", dX, dY)
	}

	if newDirection == previousDirection {
		return 1, newDirection
	} else if (previousDirection == advent.Down || previousDirection == advent.Up) && (newDirection == advent.Left || newDirection == advent.Right) {
		return 1001, newDirection
	} else if (previousDirection == advent.Left || previousDirection == advent.Right) && (newDirection == advent.Up || newDirection == advent.Down) {
		return 1001, newDirection
	} else if (previousDirection == advent.Up && newDirection == advent.Down) ||
		(previousDirection == advent.Down && newDirection == advent.Up) ||
		(previousDirection == advent.Left && newDirection == advent.Right) ||
		(previousDirection == advent.Right && newDirection == advent.Left) {
		return 2001, newDirection
	}

	log.Fatalf("unexpected direction traversal from %s to %s", previousDirection, newDirection)
	return 0, newDirection
}

type traversal struct {
	cost         int
	previousNode *advent.Point
	dir          *advent.Direction
}

const inf = math.MaxInt

var adjacencts func(pos advent.Point) []advent.Point

func getPathGraph(m [][]object, start advent.Point) map[advent.Point][]advent.Point {
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
			if m[adj.Y][adj.X] != wall {
				g[current] = append(g[current], adj)
				toVisit = append(toVisit, adj)
			}
		}
	}
	return g
}

type object int

const (
	path object = iota
	wall
	start
	end
)

// Return the map, starting, and ending position (in that order)
func mustParseInput(r io.Reader) ([][]object, advent.Point, advent.Point) {
	var m [][]object
	var startPoint, endPoint advent.Point

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.Split(s.Text(), "")
		row := make([]object, len(line))

		for i := range line {
			switch line[i] {
			case ".":
				row[i] = path
			case "#":
				row[i] = wall
			case "S":
				row[i] = start
				startPoint = advent.Point{X: i, Y: len(m)}
			case "E":
				row[i] = end
				endPoint = advent.Point{X: i, Y: len(m)}
			default:
				log.Fatalf("invalid character at (%d,%d)", i, len(m))
			}
		}

		m = append(m, row)
	}

	return m, startPoint, endPoint
}

func printMap(m [][]object, shortestPath map[advent.Point]*traversal, startPoint advent.Point, endPoint advent.Point) {
	lines := [][]string{}
	for y := range m {
		lines = append(lines, make([]string, len(m[y])))
		for x := range m[y] {
			switch m[y][x] {
			case path:
				lines[y][x] = "."
			case wall:
				lines[y][x] = "#"
			case start:
				lines[y][x] = "S"
			case end:
				lines[y][x] = "E"
			}
		}
	}

	nextNode := endPoint
	for {
		traversal := shortestPath[nextNode]
		if traversal.previousNode == nil || *traversal.previousNode == startPoint {
			break
		}
		pn := traversal.previousNode
		var dir string
		if *traversal.dir == advent.Up {
			dir = "^"
		} else if *traversal.dir == advent.Down {
			dir = "v"
		} else if *traversal.dir == advent.Left {
			dir = "<"
		} else if *traversal.dir == advent.Right {
			dir = ">"
		}
		lines[pn.Y][pn.X] = dir
		nextNode = *traversal.previousNode
	}

	for i := range lines {
		fmt.Println(strings.Join(lines[i], ""))
	}
}

func printPath(shortestPath map[advent.Point]*traversal, endPoint advent.Point) {
	type nodeTraversal struct {
		node advent.Point
		trav traversal
	}
	backwards := []nodeTraversal{}
	nextNode := endPoint
	for {
		trav := shortestPath[nextNode]
		backwards = append(backwards, nodeTraversal{node: nextNode, trav: *trav})
		if trav.previousNode == nil {
			break
		}
		nextNode = *trav.previousNode
	}

	for i := len(backwards) - 1; i >= 0; i-- {
		nt := backwards[i]
		log.Printf("%s to %v incurring total cost %d", nt.trav.dir, nt.node, nt.trav.cost)
	}
}
