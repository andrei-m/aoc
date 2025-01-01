package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	m, start, end := mustParseInput(os.Stdin)
	log.Printf("%dx%d maze; start: %v; end: %v", len(m[0]), len(m), start, end)
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
