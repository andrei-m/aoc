package advent

import "math"

type edge struct {
	cost         int
	previousNode Point
}

const inf = math.MaxInt

func GetShortestPath(pathGraph map[Point][]Point, start Point) map[Point]Point {
	distances := map[Point]edge{}
	toVisit, visited := map[Point]struct{}{}, map[Point]struct{}{}

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
		var visitNext Point
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

	result := make(map[Point]Point, len(distances))
	for k, v := range distances {
		result[k] = v.previousNode
	}
	return result
}

func TraverseShortestPath(shortestPath map[Point]Point, end Point) int {
	pos := end
	steps := 0
	for {
		if pos.X == 0 && pos.Y == 0 {
			return steps
		}
		pos = shortestPath[pos]
		steps++
	}
}
