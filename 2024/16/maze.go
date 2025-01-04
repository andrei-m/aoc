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
	path := getShortestPaths(pathGraph, start)

	//printMap(m, path, start, end)
	//printPath(path, node{loc: end, dir: advent.Right})
	fmt.Printf("part 1: %d\n", path[node{loc: end, dir: advent.Right}].cost)

	/*
		part 2: Keep track of all previous nodes with equally good costs in traversal.previousNode. Starting with 'E', walk the tree and count the distinct nodes collected this way.
	*/
}

// Build a map of destination->traversal cost/previous path step for each destination reachable from 'start' using Djikstra's algorithm
func getShortestPaths(pathGraph map[advent.Point][]node, start advent.Point) map[node]*traversal {
	debugEnabled := advent.DebugEnabled()
	distances := map[node]*traversal{}
	toVisit := map[node]struct{}{}
	visited := map[node]struct{}{}

	for _, nodes := range pathGraph {
		for _, n := range nodes {
			distances[n] = &traversal{cost: inf}
			toVisit[n] = struct{}{}
		}
	}
	// initial distances from 'start' to start, assuming initial direction is East (right)
	distances[node{loc: start, dir: advent.Right}] = &traversal{cost: 0}
	toVisit[node{loc: start, dir: advent.Right}] = struct{}{}

	// Visit each element of 'toVisit' in priority order, according to lowest cost in 'distances'
	for {
		if len(toVisit) == 0 {
			// visited all nodes
			break
		}

		//TODO: prefer a heap for more efficient priority queue
		lowestCost := inf
		var lowestCostNode node
		for n := range toVisit {
			if distances[n].cost < lowestCost {
				lowestCost = distances[n].cost
				lowestCostNode = n
			}
		}
		if debugEnabled {
			log.Printf("lowest cost node: %v", lowestCostNode)
		}

		lcnTraversal := distances[lowestCostNode]
		adjacents := pathGraph[lowestCostNode.loc]
		for i := range adjacents {
			adjacentNode := adjacents[i]
			_, previouslyVisited := visited[adjacentNode]
			if previouslyVisited {
				continue
			}

			edgeCost := cost(lowestCostNode, adjacentNode)
			fullTraveralCost := lcnTraversal.cost + edgeCost

			if fullTraveralCost < distances[adjacentNode].cost {
				previousCost := distances[adjacentNode].cost

				distances[adjacentNode].cost = fullTraveralCost
				distances[adjacentNode].previousNode = &lowestCostNode

				if debugEnabled {
					log.Printf("found lower cost %d (from %d) to get to %v from %v", distances[adjacentNode].cost, previousCost, adjacentNode, lowestCostNode)
				}
			}
		}

		delete(toVisit, lowestCostNode)
		visited[lowestCostNode] = struct{}{}
	}

	return distances
}

// Calculate the cost and new direction of a potential edge traversal
func cost(previousNode, nextNode node) int {
	dX := nextNode.loc.X - previousNode.loc.X
	dY := nextNode.loc.Y - previousNode.loc.Y
	if advent.Abs(dX)+advent.Abs(dY) != 1 {
		log.Fatalf("%v and %v are not adjacent", previousNode, nextNode)
	}
	return 1 + rotationCost(previousNode.dir, nextNode.dir)
}

func rotationCost(dirA, dirB advent.Direction) int {
	if dirA == dirB {
		return 0
	}
	if advent.RotateClockwise(dirA) == dirB || advent.RotateCounterClockwise(dirA) == dirB {
		return 1000
	}
	return 2000
}

type node struct {
	loc advent.Point
	dir advent.Direction
}

func (n node) String() string {
	return fmt.Sprintf("{%s %s}", n.loc, n.dir)
}

type traversal struct {
	cost         int
	previousNode *node
}

const inf = math.MaxInt

var adjacencts func(pos advent.Point) []advent.Point

func getPathGraph(m [][]object, start advent.Point) map[advent.Point][]node {
	g := map[advent.Point][]node{}
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
				dX := adj.X - current.X
				dY := adj.Y - current.Y

				var dir advent.Direction
				if dX == -1 {
					dir = advent.Left
				} else if dX == 1 {
					dir = advent.Right
				} else if dY == -1 {
					dir = advent.Up
				} else if dY == 1 {
					dir = advent.Down
				} else {
					log.Fatalf("unexpected dX: %d or dY: %d", dX, dY)
				}

				g[current] = append(g[current], node{loc: adj, dir: dir})
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

func printMap(m [][]object, shortestPath map[node]*traversal, startPoint advent.Point, endPoint advent.Point) {
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

	var lcnEndpoint node
	endpointCost := inf
	for _, dir := range advent.Dirs {
		potentialLCN := node{loc: endPoint, dir: dir}
		trav, exists := shortestPath[potentialLCN]
		if !exists {
			continue
		}
		if trav.cost < endpointCost {
			endpointCost = trav.cost
			lcnEndpoint = potentialLCN
		}
	}
	if endpointCost == inf {
		log.Fatalf("couldn't find path to %v", endPoint)
	}

	nextNode := lcnEndpoint
	for {
		traversal := shortestPath[nextNode]
		if traversal.previousNode == nil || traversal.previousNode.loc == startPoint {
			break
		}
		pn := traversal.previousNode
		var dir string
		if pn.dir == advent.Up {
			dir = "^"
		} else if pn.dir == advent.Down {
			dir = "v"
		} else if pn.dir == advent.Left {
			dir = "<"
		} else if pn.dir == advent.Right {
			dir = ">"
		}
		lines[pn.loc.Y][pn.loc.X] = dir
		nextNode = *traversal.previousNode
	}

	for i := range lines {
		fmt.Println(strings.Join(lines[i], ""))
	}
}

func printPath(shortestPath map[node]*traversal, endPoint node) {
	type nodeTraversal struct {
		n    node
		trav traversal
	}
	backwards := []nodeTraversal{}
	nextNode := endPoint
	for {
		trav := shortestPath[nextNode]
		backwards = append(backwards, nodeTraversal{n: nextNode, trav: *trav})
		if trav.previousNode == nil {
			break
		}
		nextNode = *trav.previousNode
	}

	for i := len(backwards) - 1; i >= 0; i-- {
		nt := backwards[i]
		log.Printf("%s to %v incurring total cost %d", nt.n.dir, nt.n.loc, nt.trav.cost)
	}
}
