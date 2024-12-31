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
	m, dirs := mustParseInput()
	log.Printf("map: %dx%d with %d directions", len(m[0]), len(m), len(dirs))
	adjacentDir = advent.AdjacentDirFn(len(m[0]), len(m))

	var robotPos advent.Point
	for y := range m {
		for x := range m[y] {
			if m[y][x] == robot {
				robotPos = advent.Point{X: x, Y: y}
			}
		}
	}

	/*
		A robot OR a box can move in the direction of an adjacenct space iff:
			- the adjacent space is empty
			- the adjacent space has a moveable box
		To move a robot or box:
			- Swap its position with the adjacent empty space
			- A sequence of movable boxes/robots should start moving (swapping) adjacent to its empty space
	*/
	for _, dir := range dirs {
		boxesMoved := 0
		for {
			vacatedPos, newPos := maybeMove(m, robotPos, dir)
			if vacatedPos == nil {
				/*
					drawMap(m)
					fmt.Printf("direction #%d: cannot move %s\n", i+1, dir)
					time.Sleep(1 * time.Second)
				*/
				// can't move in direction
				break
			}
			if *vacatedPos == robotPos {
				// robot moved
				robotPos = *newPos
				/*
					drawMap(m)
					boxMoveSuffix := ""
					if boxesMoved == 1 {
						boxMoveSuffix = " (moved 1 box)"
					} else if boxesMoved > 1 {
						boxMoveSuffix = fmt.Sprintf(" (moved %d boxes)", boxesMoved)
					}
					fmt.Printf("direction #%d: moved %s%s\n", i+1, dir, boxMoveSuffix)
					time.Sleep(1 * time.Second)
				*/
				break
			}
			// moved a box; try moving again
			boxesMoved++
		}
	}

	fmt.Printf("part 1: %d\n", score(m))
}

var adjacentDir func(advent.Point, advent.Direction) *advent.Point

func maybeMove(m [][]object, pos advent.Point, dir advent.Direction) (*advent.Point, *advent.Point) {
	adj := adjacentDir(pos, dir)
	if adj == nil {
		// reached a map edge
		return nil, nil
	}
	switch m[adj.Y][adj.X] {
	case empty:
		m[adj.Y][adj.X], m[pos.Y][pos.X] = m[pos.Y][pos.X], m[adj.Y][adj.X]
		return &pos, adj
	case wall:
		return nil, nil
	case box:
		return maybeMove(m, *adj, dir)
	default:
		log.Fatalf("unexpected object at position (%d,%d)", adj.X, adj.Y)
	}
	return nil, nil
}

func score(m [][]object) int {
	sum := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] == box {
				sum += 100*y + x
			}
		}
	}
	return sum
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

func drawMap(m [][]object) {
	for y := range m {
		sb := strings.Builder{}
		for x := range m[y] {
			switch m[y][x] {
			case empty:
				sb.WriteString(".")
			case box:
				sb.WriteString("O")
			case wall:
				sb.WriteString("#")
			case robot:
				sb.WriteString("@")
			}
		}
		fmt.Println(sb.String())
	}
}
