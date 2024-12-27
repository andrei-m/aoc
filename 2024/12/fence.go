package main

import (
	"bufio"
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
				reg = regions[adj]
				if reg != nil && reg.plant == rows[y][x] {
					// found region to join
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

	log.Printf("region at (70,16): %v", *regions[advent.Point{X: 70, Y: 16}])
}

type region struct {
	plant string
	locs  []advent.Point
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
