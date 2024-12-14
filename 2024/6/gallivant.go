package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andrei-m/aoc/advent"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

func main() {
	debug := advent.DebugEnabled()
	dir := up
	var guardPosition advent.Point
	obstacles := map[advent.Point]struct{}{}

	y := 0
	xOverflow := -1
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x, chr := range line {
			if chr == "^" {
				guardPosition = advent.Point{X: x, Y: y}
			} else if chr == "#" {
				obstacles[advent.Point{X: x, Y: y}] = struct{}{}
			}
		}

		y++
		if xOverflow == -1 {
			xOverflow = len(line)
		}
	}

	// include the starting position
	pointsVisited := map[advent.Point][]direction{
		guardPosition: {dir},
	}
	addedObstacleLocs := map[advent.Point]struct{}{}
	for {
		move := nextMove(guardPosition, dir)
		if move.X < 0 || move.X == xOverflow || move.Y < 0 || move.Y == y {
			// exited the area; do not count the point
			break
		}

		_, hitObstacle := obstacles[move]
		if hitObstacle {
			dir = nextDirection(dir)
			continue
		}
		// move is OK

		previousDirs, ok := pointsVisited[move]
		if ok {
			// Point was visited previously navigating in the **next** direction, so it's a loop candidate
			loopDirection := nextDirection(dir)
			if dirSliceContains(previousDirs, loopDirection) {
				obstacleLoc := nextMove(move, dir)
				if !(obstacleLoc.X < 0 || obstacleLoc.X == xOverflow || obstacleLoc.Y < 0 || obstacleLoc.Y == y) {
					addedObstacleLocs[obstacleLoc] = struct{}{}
				}
			}
		}

		// record the move
		pointsVisited[move] = append(pointsVisited[move], dir)
		guardPosition = move

		if debug {
			drawMap(guardPosition, pointsVisited, xOverflow, y, obstacles, addedObstacleLocs)
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Printf("part 1: %d\n", len(pointsVisited))
	fmt.Printf("part 2: %d\n", len(addedObstacleLocs))
}

func drawMap(pos advent.Point, pointsVisited map[advent.Point][]direction, xOverflow int, yOverflow int, obstacles map[advent.Point]struct{}, addedObstacleLocs map[advent.Point]struct{}) {
	for y := 0; y < yOverflow; y++ {
		for x := 0; x < xOverflow; x++ {
			p := advent.Point{X: x, Y: y}
			if pos == p {
				fmt.Print("X")
				continue
			}

			_, addedObstacle := addedObstacleLocs[p]
			if addedObstacle {
				fmt.Print("O")
				continue
			}
			_, pointVisted := pointsVisited[p]
			if pointVisted {
				fmt.Print("x")
				continue
			}
			_, obstacle := obstacles[p]
			if obstacle {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
	fmt.Printf("pos: %s, points visited: %d, added obstacles candidates: %d", pos, len(pointsVisited), len(addedObstacleLocs))
}

func nextMove(pos advent.Point, dir direction) advent.Point {
	switch dir {
	case up:
		return advent.Point{X: pos.X, Y: pos.Y - 1}
	case right:
		return advent.Point{X: pos.X + 1, Y: pos.Y}
	case down:
		return advent.Point{X: pos.X, Y: pos.Y + 1}
	case left:
		return advent.Point{X: pos.X - 1, Y: pos.Y}
	}
	panic("invalid direction")
}

func nextDirection(dir direction) direction {
	switch dir {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	panic("invalid direction")
}

func dirSliceContains(sl []direction, val direction) bool {
	for i := range sl {
		if sl[i] == val {
			return true
		}
	}
	return false
}
