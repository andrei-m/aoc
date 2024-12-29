package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	debug := advent.DebugEnabled()
	robots := mustParseRobots()
	log.Printf("parsed %d robots", len(robots))

	for i := 0; i < duration; i++ {
		for j := range robots {
			robots[j].advance()
			if debug && j == 0 {
				log.Printf("robot index %d after %d seconds: %v", j, i+1, robots[j])
			}
		}
	}

	fmt.Printf("part 1: %d\n", safetyFactor(robots))
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
