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
	debug := advent.DebugEnabled()
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

	part1Antinodes := map[advent.Point]struct{}{}
	part2Antinodes := map[advent.Point]struct{}{}

	for freq, ants := range antennas {
		p := pairs(ants)

		for _, pair := range p {
			inverted := advent.InvertVector(advent.Vector(pair))
			if inverted.B.Inbounds(xOverflow, yOverflow) {
				part1Antinodes[inverted.B] = struct{}{}
			}
		}

		if debug {
			log.Printf("%s: %v", freq, ants)
			log.Printf("%v", p)
			log.Println()
		}

		for _, pair := range p {
			v := advent.Vector(pair)
			step := shortenInline(v)
			for {
				if v.B.Inbounds(xOverflow, yOverflow) {
					part2Antinodes[v.B] = struct{}{}
				} else {
					break
				}
				v = advent.AddVector(v, step)
			}
			// The inverse direction is handled by the inverse pair, which is present in 'pairs'
		}
	}

	fmt.Printf("part 1: %d\n", len(part1Antinodes))
	fmt.Printf("part 2: %d\n", len(part2Antinodes))
}

func shortenInline(v advent.Vector) advent.Vector {
	dX := v.A.X - v.B.X
	dY := v.A.Y - v.B.Y

	absDX := advent.Abs(dX)
	absDY := advent.Abs(dY)
	gcd := advent.GCD(absDX, absDY)

	return advent.Vector{
		A: v.A,
		B: advent.Point{X: v.A.X + dX/gcd, Y: v.A.Y + dY/gcd},
	}
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
