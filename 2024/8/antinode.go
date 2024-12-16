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

	antinodes := map[advent.Point]struct{}{}

	for freq, ants := range antennas {
		p := pairs(ants)
		for _, pair := range p {
			inverted := advent.InvertVector(advent.Vector(pair))
			if inverted.B.Inbounds(xOverflow, yOverflow) {
				antinodes[inverted.B] = struct{}{}
			}
		}
		log.Printf("%s: %v", freq, ants)
		log.Printf("%v", p)
		log.Println()
	}
	fmt.Printf("part 1: %d\n", len(antinodes))

	/*
			part 2 outline:
			For each pair's vector, determine the shortest vector that is still inline
			- This probably means finding the greatest common divisor between dX & dY, then creating a shorter vector after dividing by that amount
		 	iterate in increments of that vector, accumulating each point until overflow
			do the same with the inverse of that vector
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
