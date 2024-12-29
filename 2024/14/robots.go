package main

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	robots := mustParseRobots()
	for _, r := range robots {
		log.Printf("robot: %v", r)
	}
	log.Printf("parsed %d robots", len(robots))
}

const (
	width    = 101
	height   = 103
	duration = 100
)

type robot struct {
	pos advent.Point
	vel advent.Point
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
