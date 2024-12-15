package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	xOverflow := -1
	yOverflow := 0
	scanner := bufio.NewScanner(os.Stdin)

	antennas := map[string][]advent.Point{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x, chr := range line {
			if chr != "." {
				antennas[chr] = append(antennas[chr], advent.Point{X: x, Y: yOverflow})
			}
		}

		yOverflow++
		if xOverflow == -1 {
			xOverflow = len(line)
		}
	}

	for freq, ants := range antennas {
		p := pairs(ants)
		log.Printf("%s: %v", freq, ants)
		log.Printf("%v", p)
		log.Println()
	}
	/*
		For each frequency (antennas key):
		enumerate pairs of antennas
		for each pair (A, B):
		evaluate vectors A->B and B->A
		invert each vector
		evaluate the new point relative to origin A or B
		if not overflowing the map, count the resulting point
	*/
}

func pairs(points []advent.Point) []advent.Pair[advent.Point] {
	if len(points) < 2 {
		return []advent.Pair[advent.Point]{}
	}

	result := []advent.Pair[advent.Point]{}
	for i := range points {
		for j := range points {
			if i == j {
				continue
			}
			result = append(result, advent.Pair[advent.Point]{A: points[i], B: points[j]})
		}
	}

	return result
}
