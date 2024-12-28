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
	log.Printf("x Overflow: %d; yOverflow: %d", xOverflow, yOverflow)

	regions := map[advent.Point]*region{}
	for y := range rows {
		for x := range rows[y] {
			loc := advent.Point{X: x, Y: y}
			var reg *region

			for _, adj := range adjacents(loc, xOverflow, yOverflow) {
				testReg := regions[adj]
				if testReg != nil && testReg.plant == rows[y][x] {
					//TODO: make this account for multiple regions that should be merged
					// found region to join
					reg = testReg
					break
				}
			}

			if reg != nil {
				reg.locs = append(reg.locs, loc)
			} else {
				// new region
				reg = &region{plant: rows[y][x], locs: []advent.Point{loc}}
			}
			regions[loc] = reg
		}
	}

	dedupedRegions := map[*region]struct{}{}
	for _, region := range regions {
		reg := region
		dedupedRegions[reg] = struct{}{}
	}

	var score int
	for region := range dedupedRegions {
		score += scoreRegion(*region, xOverflow, yOverflow)
	}
	fmt.Printf("%d regions part 1: %d\n", len(dedupedRegions), score)

	log.Printf("region at (70,16): %v", *regions[advent.Point{X: 131, Y: 131}])
	scoreRegion(*regions[advent.Point{X: 131, Y: 131}], xOverflow, yOverflow)
}

type region struct {
	plant string
	locs  []advent.Point
}

func scoreRegion(r region, xOverflow int, yOverflow int) int {
	area := len(r.locs)

	locLookup := make(map[advent.Point]struct{}, area)
	for _, loc := range r.locs {
		locLookup[loc] = struct{}{}
	}

	var perimeter int
	for _, loc := range r.locs {
		adjs := adjacents(loc, xOverflow, yOverflow)
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

	log.Printf("region %v: area: %d, perimeter: %d", r, area, perimeter)
	return area * perimeter
}

func adjacents(loc advent.Point, xOverflow int, yOverflow int) []advent.Point {
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

func parsePlantRows() [][]string {
	rows := [][]string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), ""))
	}
	return rows
}
