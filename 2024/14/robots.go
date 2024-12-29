package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	debug := advent.DebugEnabled()
	robots := mustParseRobots()
	log.Printf("parsed %d robots", len(robots))

	elapsed := 0
	for i := 0; i < duration; i++ {
		elapsed++
		for j := range robots {
			robots[j].advance()
			if debug && j == 0 {
				log.Printf("robot index %d after %d seconds: %v", j, i+1, robots[j])
			}
		}
	}

	fmt.Printf("part 1: %d\n", safetyFactor(robots))

	//TODO: account for the possibility of the Christmas tree in the first 'duration' iterations. Currently, it is assumed to show up after 'duration'
	for {
		elapsed++
		for i := range robots {
			robots[i].advance()
		}

		positions := map[advent.Point]struct{}{}
		for i := range robots {
			positions[robots[i].pos] = struct{}{}
		}
		maxLength := 0
		for pos := range positions {
			cl := chainLength(positions, pos)
			if cl > maxLength {
				maxLength = cl
			}
		}

		if maxLength > 200 {
			draw(robots)
			fmt.Printf("after %d seconds; longest chain: %d", elapsed, maxLength)
			fmt.Println()
			time.Sleep(5 * time.Second)
		}
	}
}

func chainLength(robotPositions map[advent.Point]struct{}, pos advent.Point) int {
	visited := map[advent.Point]struct{}{}
	return searchLength(robotPositions, pos, visited)
}

func searchLength(robotPositions map[advent.Point]struct{}, pos advent.Point, visited map[advent.Point]struct{}) int {
	_, robotPresent := robotPositions[pos]
	if !robotPresent {
		return 0
	}
	_, previouslyVisited := visited[pos]
	if previouslyVisited {
		return 0
	}
	visited[pos] = struct{}{}
	chainLength := 1
	for _, adj := range adjacents(pos) {
		chainLength += searchLength(robotPositions, adj, visited)
	}
	return chainLength
}

func adjacents(pos advent.Point) []advent.Point {
	/*
		horizontal, vertical, and diagonal adjacents ('A') relative to position ('P')
		AAA
		APA
		AAA
	*/
	adj := []advent.Point{}
	if pos.X > 0 {
		// left
		adj = append(adj, advent.Point{X: pos.X - 1, Y: pos.Y})
	}
	if pos.X < width-1 {
		// right
		adj = append(adj, advent.Point{X: pos.X + 1, Y: pos.Y})
	}
	if pos.Y > 0 {
		// up
		adj = append(adj, advent.Point{X: pos.X, Y: pos.Y - 1})
	}
	if pos.Y < height-1 {
		// down
		adj = append(adj, advent.Point{X: pos.X, Y: pos.Y + 1})
	}

	if pos.X > 0 && pos.Y > 0 {
		// up and left
		adj = append(adj, advent.Point{X: pos.X - 1, Y: pos.Y - 1})
	}
	if pos.X < width-1 && pos.Y < height-1 {
		// down and right
		adj = append(adj, advent.Point{X: pos.X + 1, Y: pos.Y + 1})
	}
	if pos.X > 0 && pos.Y < height-1 {
		// down and left
		adj = append(adj, advent.Point{X: pos.X - 1, Y: pos.Y + 1})
	}
	if pos.X < width-1 && pos.Y > 0 {
		// up and right
		adj = append(adj, advent.Point{X: pos.X + 1, Y: pos.Y - 1})
	}

	return adj
}

func safetyFactor(robots []robot) int {
	quadrents := make([]int, 4)
	for _, r := range robots {
		if r.pos.X < width/2 && r.pos.Y < height/2 {
			quadrents[0]++
		} else if r.pos.X > width/2 && r.pos.Y < height/2 {
			quadrents[1]++
		} else if r.pos.X > width/2 && r.pos.Y > height/2 {
			quadrents[2]++
		} else if r.pos.X < width/2 && r.pos.Y > height/2 {
			quadrents[3]++
		}
	}
	log.Printf("quadrents: %v", quadrents)
	return quadrents[0] * quadrents[1] * quadrents[2] * quadrents[3]
}

var (
	// make variable for testing purposes
	width  = 101
	height = 103
	/*
		quadrent 0: X < 50 && Y < 51
		quadrent 1: X > 50 && Y < 51
		quadrent 2: X > 50 && Y > 51
		quadrent 3: X < 50 && Y > 51
	*/
	duration = 100
)

type robot struct {
	pos advent.Point
	vel advent.Point
}

func (r *robot) advance() {
	nextPos := advent.Point{X: r.pos.X + r.vel.X, Y: r.pos.Y + r.vel.Y}
	if nextPos.X < 0 {
		nextPos.X = width + nextPos.X
	}
	if nextPos.X >= width {
		nextPos.X = nextPos.X - width
	}
	if nextPos.Y < 0 {
		nextPos.Y = height + nextPos.Y
	}
	if nextPos.Y >= height {
		nextPos.Y = nextPos.Y - height
	}
	r.pos = nextPos
}

func draw(robots []robot) {
	positions := map[advent.Point]struct{}{}
	for i := range robots {
		positions[robots[i].pos] = struct{}{}
	}
	for y := 0; y < height; y++ {
		sb := strings.Builder{}
		for x := 0; x < width; x++ {
			_, hasRobot := positions[advent.Point{X: x, Y: y}]
			if hasRobot {
				sb.WriteString("R")
			} else {
				sb.WriteString(".")
			}
		}
		fmt.Println(sb.String())
	}
}

var (
	robotPattern = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
)

func mustParseRobots() []robot {
	robots := []robot{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		submatch := robotPattern.FindStringSubmatch(line)
		if len(submatch) != 5 {
			log.Fatalf("invalid line: %s", line)
		}

		pos := advent.Point{X: advent.MustParseInt(submatch[1]), Y: advent.MustParseInt(submatch[2])}
		vel := advent.Point{X: advent.MustParseInt(submatch[3]), Y: advent.MustParseInt(submatch[4])}
		robots = append(robots, robot{pos: pos, vel: vel})
	}

	return robots
}
