package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	m, dirs := mustParseInput(os.Stdin)
	originalM := make([][]object, len(m))
	for y := range m {
		row := make([]object, len(m[y]))
		copy(row, m[y])
		originalM[y] = row
	}

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

	/*
		part2:
		A robot can move into an adjacent space iff:
			- the adjacent space is empty
			- the adjacent space has a movable box
		A box can move left or right iff:
			- the adjacent space is empty
			- the adjacent space has a movable box
		A box can move up or down iff:
			- both adjacent spaces are empty or contain movable boxes
		left and right halves of boxes must move together
	*/
	part2(originalM, dirs)
}

func part2(m [][]object, dirs []advent.Direction) {
	mPart2 := getPart2Map(m)
	adjacentDir = advent.AdjacentDirFn(len(mPart2[0]), len(mPart2))

	var robotPos advent.Point
	for y := range mPart2 {
		for x := range mPart2[y] {
			if mPart2[y][x] == robot {
				robotPos = advent.Point{X: x, Y: y}
			}
		}
	}

	for i, dir := range dirs {
		var nextMoveSuffix string
		if i < len(dirs)-1 {
			nextMoveSuffix = fmt.Sprintf("; next move: %v", dirs[i+1])
		}
		boxesMoved := 0
		for {
			vacatedPos, newPos := maybeMove(mPart2, robotPos, dir)
			if vacatedPos == nil {
				drawMap(mPart2)
				fmt.Printf("(%d,%d) direction #%d: cannot move %s%s\n", robotPos.X, robotPos.Y, i+1, dir, nextMoveSuffix)
				time.Sleep(1 * time.Second)
				// can't move in direction
				break
			}
			if *vacatedPos == robotPos {
				// robot moved
				robotPos = *newPos
				drawMap(mPart2)

				boxMoveSuffix := ""
				if boxesMoved == 1 {
					boxMoveSuffix = " (moved 1 box)"
				} else if boxesMoved > 1 {
					boxMoveSuffix = fmt.Sprintf(" (moved %d boxes)", boxesMoved)
				}
				fmt.Printf("(%d,%d) direction #%d: moved %s%s%s\n", robotPos.X, robotPos.Y, i+1, dir, boxMoveSuffix, nextMoveSuffix)

				time.Sleep(1 * time.Second)
				break
			}
			// moved a box; try moving again
			boxesMoved++
		}
	}

	fmt.Printf("part 2: %d\n", score(mPart2))
}

var adjacentDir func(advent.Point, advent.Direction) *advent.Point

// maybe move returns the vacated and new position of the moved object, or nil if no move occured
// For 2x1 boxes, the left half is represented in returned Points (both halves move)
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
	case box, boxLeft, boxRight:
		return maybeMoveBox(m, *adj, dir)
	default:
		log.Fatalf("unexpected object at position (%d,%d)", adj.X, adj.Y)
	}
	return nil, nil
}

func maybeMoveBox(m [][]object, pos advent.Point, dir advent.Direction) (*advent.Point, *advent.Point) {
	if m[pos.Y][pos.X] == box {
		// 1x1 boxes use the same movement rules as robots
		return maybeMove(m, pos, dir)
	}

	var posLeft, posRight advent.Point
	if m[pos.Y][pos.X] == boxLeft {
		posLeft = pos
		posRight = *adjacentDir(pos, advent.Right)
	} else if m[pos.Y][pos.X] == boxRight {
		posLeft = *adjacentDir(pos, advent.Left)
		posRight = pos
	} else {
		log.Fatalf("(%d,%d) does not contain a box", pos.X, pos.Y)
	}

	adjLeft := adjacentDir(posLeft, dir)
	adjRight := adjacentDir(posRight, dir)
	switch dir {
	case advent.Down, advent.Up:
		if adjLeft == nil || adjRight == nil {
			// reached a map edge
			return nil, nil
		}
		if m[adjLeft.Y][adjLeft.X] == wall || m[adjRight.Y][adjRight.X] == wall {
			// wall adjacent to one of the halves
			return nil, nil
		}
		if m[adjLeft.Y][adjLeft.X] == boxLeft || m[adjLeft.Y][adjLeft.X] == boxRight || m[adjLeft.Y][adjLeft.X] == box {
			return maybeMoveBox(m, *adjLeft, dir)
		}
		if m[adjRight.Y][adjRight.X] == boxRight || m[adjRight.Y][adjRight.X] == boxLeft || m[adjRight.Y][adjRight.X] == box {
			return maybeMoveBox(m, *adjRight, dir)
		}
		// Neither adjacency contains a wall, box, or edge, so they are empty
		m[posLeft.Y][posLeft.X], m[adjLeft.Y][adjLeft.X] = m[adjLeft.Y][adjLeft.X], m[posLeft.Y][posLeft.X]
		m[posRight.Y][posRight.X], m[adjRight.Y][adjRight.X] = m[adjRight.Y][adjRight.X], m[posRight.Y][posRight.X]
		return &pos, adjLeft
	case advent.Left:
		if adjLeft == nil {
			return nil, nil
		}
		if m[adjLeft.Y][adjLeft.X] == wall {
			return nil, nil
		}
		if m[adjLeft.Y][adjLeft.X] == boxRight || m[adjLeft.Y][adjLeft.X] == box {
			return maybeMoveBox(m, *adjLeft, dir)
		}
		// move the left side first to not overwrite the left side of the box with the right side
		// left moves over one to the adjacent spot; right swaps with the spot vacated by left
		m[posLeft.Y][posLeft.X], m[adjLeft.Y][adjLeft.X] = m[adjLeft.Y][adjLeft.X], m[posLeft.Y][posLeft.X]
		m[posRight.Y][posRight.X], m[posLeft.Y][posLeft.X] = m[posLeft.Y][posLeft.X], m[posRight.Y][posRight.X]
		return &posLeft, adjLeft
	case advent.Right:
		if adjRight == nil {
			return nil, nil
		}
		if m[adjRight.Y][adjRight.X] == wall {
			return nil, nil
		}
		if m[adjRight.Y][adjRight.X] == boxLeft || m[adjRight.Y][adjRight.X] == box {
			return maybeMoveBox(m, *adjRight, dir)
		}
		// move the right side into the adjacency; swap the left into the spot vacated by right
		m[posRight.Y][posRight.X], m[adjRight.Y][adjRight.X] = m[adjRight.Y][adjRight.X], m[posRight.Y][posRight.X]
		m[posLeft.Y][posLeft.X], m[posRight.Y][posRight.X] = m[posRight.Y][posRight.X], m[posLeft.Y][posLeft.X]
		return &posLeft, &posRight
	}

	return nil, nil
}

func score(m [][]object) int {
	sum := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] == box || m[y][x] == boxLeft {
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
	boxLeft
	boxRight
	robot
)

func mustParseInput(r io.Reader) ([][]object, []advent.Direction) {
	scanner := bufio.NewScanner(r)
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
			case boxLeft:
				sb.WriteString("[")
			case boxRight:
				sb.WriteString("]")
			case wall:
				sb.WriteString("#")
			case robot:
				sb.WriteString("@")
			}
		}
		fmt.Println(sb.String())
	}
}

func getPart2Map(m [][]object) [][]object {
	result := make([][]object, len(m))
	for y := range m {
		row := make([]object, len(m[y])*2)
		for x := range m[y] {
			switch m[y][x] {
			case empty:
				row[x*2] = empty
				row[x*2+1] = empty
			case box:
				row[x*2] = boxLeft
				row[x*2+1] = boxRight
			case wall:
				row[x*2] = wall
				row[x*2+1] = wall
			case robot:
				row[x*2] = robot
				row[x*2+1] = empty
			default:
				log.Fatalf("unexpected object at (%d,%d)", x, y)
			}
		}
		result[y] = row
	}
	return result
}
