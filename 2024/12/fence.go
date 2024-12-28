package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	//parse into (x, y) -> plant mappings
	// iterate over each point. A point is either its own region or merges into an existing region
	// to check for a mergable region, look for the same plant above or to the left of this point. If found, update that region to include this point

	rows := parsePlantRows()

	xOverflow := len(rows[0])
	yOverflow := len(rows)
	adjacents = adjacentsFn(xOverflow, yOverflow)
	adjacentDir = adjacentDirFn(xOverflow, yOverflow)
	log.Printf("x Overflow: %d; yOverflow: %d", xOverflow, yOverflow)

	regions := map[advent.Point]*region{}
	for y := range rows {
		for x := range rows[y] {
			loc := advent.Point{X: x, Y: y}
			_, ok := regions[loc]
			if ok {
				// location has been discovered in a region search starting from an earlier loc
				continue
			}
			reg := getRegion(loc, rows[y][x], rows)
			for i := range reg.locs {
				regions[reg.locs[i]] = &reg
			}
		}
	}

	dedupedRegions := map[*region]struct{}{}
	for _, region := range regions {
		reg := region
		dedupedRegions[reg] = struct{}{}
	}

	var score int
	for region := range dedupedRegions {
		score += scoreRegionPart1(*region)
	}
	fmt.Printf("%d regions part 1: %d\n", len(dedupedRegions), score)

	score = 0
	for region := range dedupedRegions {
		score += scoreRegionPart2(*region)
	}
	fmt.Printf("%d regions part 2: %d\n", len(dedupedRegions), score)

	reg := regions[advent.Point{X: 43, Y: 0}]
	scoreRegionPart2(*reg)
}

type region struct {
	plant string
	locs  []advent.Point
}

var (
	adjacents   func(advent.Point) []advent.Point
	adjacentDir func(advent.Point, advent.Direction) *advent.Point
)

func getRegion(loc advent.Point, plant string, plantRows [][]string) region {
	regionLocs := map[advent.Point]struct{}{}
	searchRegion(loc, plant, plantRows, regionLocs)

	locs := make([]advent.Point, 0, len(regionLocs))
	for k := range regionLocs {
		locs = append(locs, k)
	}
	return region{
		plant: plant,
		locs:  locs,
	}
}

func searchRegion(loc advent.Point, plant string, plantRows [][]string, regionLocs map[advent.Point]struct{}) {
	if plantRows[loc.Y][loc.X] != plant {
		return
	}
	regionLocs[loc] = struct{}{}
	for _, adj := range adjacents(loc) {
		// avoid infinite recursion when adjacents visit each other
		_, visited := regionLocs[adj]
		if !visited {
			searchRegion(adj, plant, plantRows, regionLocs)
		}
	}
}

func scoreRegionPart1(r region) int {
	area := len(r.locs)

	locLookup := make(map[advent.Point]struct{}, area)
	for _, loc := range r.locs {
		locLookup[loc] = struct{}{}
	}

	var perimeter int
	for _, loc := range r.locs {
		adjs := adjacents(loc)
		if len(adjs) < 4 {
			// include edges of border locations in the perimeter
			perimeter += 4 - len(adjs)
		}
		for _, adj := range adjs {
			_, ok := locLookup[adj]
			if !ok {
				perimeter += 1
			}
		}
	}

	log.Printf("P1: region %v: area: %d, perimeter: %d", r, area, perimeter)
	return area * perimeter
}

func scoreRegionPart2(r region) int {
	area := len(r.locs)
	perimeter := perimeterPart2(r)
	log.Printf("P2: region %v: area: %d, perimeter: %d", r, area, perimeter)
	return area * perimeter
}

type perimiterVisited struct {
	loc advent.Point
	dir advent.Direction
}

func perimeterPart2(r region) int {
	locLookup := make(map[advent.Point]struct{}, len(r.locs))
	for _, loc := range r.locs {
		locLookup[loc] = struct{}{}
	}
	visited := map[perimiterVisited]struct{}{}

	var perimeter int

	for _, loc := range r.locs {
		for _, dir := range advent.Dirs {
			key := perimiterVisited{
				loc: loc,
				dir: dir,
			}
			_, previouslyVisited := visited[key]
			if previouslyVisited {
				// This (loc, dir) pair is accounted for in a previous "side" visit
				continue
			}

			adj := adjacentDir(loc, dir)
			if adj != nil {
				_, plantAdjacent := locLookup[*adj]
				if plantAdjacent {
					continue
				}
				perimeter += 1
			} else {
				// border case; no plant can be adjacent
				perimeter += 1
			}
			//log.Printf("perimeter +=1: %v %s", loc, dir)

			if dir == advent.Up || dir == advent.Down {
				perimeterPart2Visit(visited, locLookup, loc, advent.Left, dir)
				perimeterPart2Visit(visited, locLookup, loc, advent.Right, dir)
			} else if dir == advent.Left || dir == advent.Right {
				perimeterPart2Visit(visited, locLookup, loc, advent.Up, dir)
				perimeterPart2Visit(visited, locLookup, loc, advent.Down, dir)
			}
		}
	}

	return perimeter
}

func perimeterPart2Visit(visited map[perimiterVisited]struct{}, locLookup map[advent.Point]struct{}, loc advent.Point, dir advent.Direction, perimiterDir advent.Direction) {
	visited[perimiterVisited{loc: loc, dir: perimiterDir}] = struct{}{}
	next := adjacentDir(loc, dir)
	if next == nil {
		return
	}
	_, nextInRegion := locLookup[*next]
	if !nextInRegion {
		return
	}
	nextPerimSide := adjacentDir(*next, perimiterDir)
	if nextPerimSide != nil {
		// Before continuing down the line, check if shape extends toward the perimeter side, which indicates a bend & the end of this visit direction
		_, nextPerimSideInRegion := locLookup[*nextPerimSide]
		if nextPerimSideInRegion {
			return
		}
	}
	perimeterPart2Visit(visited, locLookup, *next, dir, perimiterDir)
}

func adjacentsFn(xOverflow, yOverflow int) func(advent.Point) []advent.Point {
	return func(loc advent.Point) []advent.Point {
		adj := []advent.Point{}
		if loc.X > 0 {
			// left
			adj = append(adj, advent.Point{X: loc.X - 1, Y: loc.Y})
		}
		if loc.X < xOverflow-1 {
			// right
			adj = append(adj, advent.Point{X: loc.X + 1, Y: loc.Y})
		}
		if loc.Y > 0 {
			// up
			adj = append(adj, advent.Point{X: loc.X, Y: loc.Y - 1})
		}
		if loc.Y < yOverflow-1 {
			// down
			adj = append(adj, advent.Point{X: loc.X, Y: loc.Y + 1})
		}
		return adj
	}
}

func adjacentDirFn(xOverflow, yOverflow int) func(advent.Point, advent.Direction) *advent.Point {
	return func(loc advent.Point, dir advent.Direction) *advent.Point {
		switch dir {
		case advent.Left:
			if loc.X <= 0 {
				return nil
			}
			return &advent.Point{X: loc.X - 1, Y: loc.Y}
		case advent.Right:
			if loc.X >= xOverflow {
				return nil
			}
			return &advent.Point{X: loc.X + 1, Y: loc.Y}
		case advent.Up:
			if loc.Y <= 0 {
				return nil
			}
			return &advent.Point{X: loc.X, Y: loc.Y - 1}
		case advent.Down:
			if loc.Y >= yOverflow {
				return nil
			}
			return &advent.Point{X: loc.X, Y: loc.Y + 1}
		default:
			panic("invalid direction")
		}
	}
}

func parsePlantRows() [][]string {
	rows := [][]string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), ""))
	}
	return rows
}
