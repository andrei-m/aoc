package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	rotations := mustParseInput(os.Stdin)
	log.Printf("parsed %d rotations", len(rotations))
	pos := 50
	zeroCount := 0
	zeroClickCount := 0

	for _, rot := range rotations {
		originalPos := pos

		newPos, zeroPosCount := rotate(pos, rot)
		pos = newPos
		if pos == 0 {
			zeroCount++
		}
		zeroClickCount += zeroPosCount
		log.Printf("%d (%d)-> %d | zeroCount: %d | zeroClickCount: %d", originalPos, rot, pos, zeroCount, zeroClickCount)
	}
	fmt.Printf("part 1: %d\n", zeroCount)
	fmt.Printf("part 2: %d\n", zeroClickCount)
}

func rotate(pos int, rot rotation) (int, int) {
	if rot == 0 {
		return pos, 0
	}

	zeroClicks := advent.Abs(int(rot) / 100)
	effectiveRotation := rot % 100
	newPos := pos + int(effectiveRotation)
	if newPos < 0 {
		newPos = 100 + newPos
		if pos != 0 {
			zeroClicks++
		}
		return newPos, int(zeroClicks)
	} else if newPos > 99 {
		newPos = newPos - 100
		if pos != 0 {
			zeroClicks++
		}
		return newPos, int(zeroClicks)
	}

	if newPos == 0 {
		zeroClicks++
	}
	return newPos, int(zeroClicks)
}

// negative values are left; positive are right
type rotation int

func mustParseInput(r io.Reader) []rotation {
	scanner := bufio.NewScanner(r)
	rotations := []rotation{}
	for scanner.Scan() {
		line := scanner.Text()
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Failed to parse rotation distance %q: %v", line[1:], err)
		}
		if string(line[0]) == "L" {
			rotations = append(rotations, rotation(-1*dist))
		} else if string(line[0]) == "R" {
			rotations = append(rotations, rotation(dist))
		} else {
			log.Fatalf("Unexpected direction: %s", string(line[0]))
		}
	}
	return rotations
}
