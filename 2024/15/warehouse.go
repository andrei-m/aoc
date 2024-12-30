package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	m, dirs := mustParseInput()
	log.Printf("map: %dx%d with %d directions", len(m[0]), len(m), len(dirs))
}

type object int

const (
	empty object = iota
	wall
	box
	robot
)

func mustParseInput() ([][]object, []advent.Direction) {
	scanner := bufio.NewScanner(os.Stdin)
	m := mustParseMap(scanner)
	dirs := mustParseDirs(scanner)
	return m, dirs
}

func mustParseMap(scanner *bufio.Scanner) [][]object {
	m := [][]object{}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// empty line between the map & directions
			break
		}
		elements := strings.Split(line, "")
		row := make([]object, len(elements))
		for i, e := range elements {
			switch e {
			case ".":
				row[i] = empty
			case "#":
				row[i] = wall
			case "O":
				row[i] = box
			case "@":
				row[i] = robot
			default:
				log.Fatalf("invalid map character at position (%d,%d)", i, y)
			}
		}
		m = append(m, row)
		y++
	}

	return m
}

func mustParseDirs(scanner *bufio.Scanner) []advent.Direction {
	dirs := []advent.Direction{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, "")
		dirsBatch := make([]advent.Direction, len(elements))
		for i, e := range elements {
			switch e {
			case "<":
				dirsBatch[i] = advent.Left
			case "^":
				dirsBatch[i] = advent.Up
			case "v":
				dirsBatch[i] = advent.Down
			case ">":
				dirsBatch[i] = advent.Right
			default:
				log.Fatalf("invalid direction at position (%d,%d)", i, y)
			}
		}
		dirs = append(dirs, dirsBatch...)
		y++
	}
	return dirs
}
