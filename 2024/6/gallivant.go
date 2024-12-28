package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	debug := advent.DebugEnabled()
	dir := advent.Up
	var guardPosition advent.Point
	obstacles := map[advent.Point]struct{}{}

	yOverflow := 0
	xOverflow := -1
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x, chr := range line {
			if chr == "^" {
				guardPosition = advent.Point{X: x, Y: yOverflow}
			} else if chr == "#" {
				obstacles[advent.Point{X: x, Y: yOverflow}] = struct{}{}
			}
		}

		yOverflow++
		if xOverflow == -1 {
			xOverflow = len(line)
		}
	}

	originalPosition := guardPosition
	originalDirection := dir

	// include the starting position
	pointsVisited := map[advent.Point][]advent.Direction{
		guardPosition: {dir},
	}
	for {
		move := nextMove(guardPosition, dir)
		if !inBounds(move, xOverflow, yOverflow) {
			// exited the area; do not count the point
			break
		}

		_, hitObstacle := obstacles[move]
		if hitObstacle {
			dir = nextDirection(dir)
			continue
		}
		// next move does not hit an obstacle

		// record the move
		pointsVisited[move] = append(pointsVisited[move], dir)
		guardPosition = move

		if debug {
			drawMap(guardPosition, pointsVisited, xOverflow, yOverflow, obstacles)
			time.Sleep(200 * time.Millisecond)
		}
	}

	fmt.Printf("part 1: %d\n", len(pointsVisited))

	addedObstacleLocs := map[advent.Point]struct{}{}
	for y := 0; y < yOverflow; y++ {
		for x := 0; x < xOverflow; x++ {
			maybeObstacle := advent.Point{X: x, Y: y}
			_, existingObstacle := obstacles[maybeObstacle]
			if existingObstacle {
				continue
			}
			if createsLoop(originalPosition, maybeObstacle, obstacles, originalDirection, xOverflow, yOverflow) {
				addedObstacleLocs[maybeObstacle] = struct{}{}
			}
		}
	}
	fmt.Printf("part 2: %d\n", len(addedObstacleLocs))
}

func createsLoop(guardPosition advent.Point, newObstacle advent.Point, obstacles map[advent.Point]struct{}, dir advent.Direction, xOverflow int, yOverflow int) bool {
	pointsVisited := map[advent.Point][]advent.Direction{
		guardPosition: {dir},
	}
	for {
		move := nextMove(guardPosition, dir)
		if !inBounds(move, xOverflow, yOverflow) {
			// exited the area; do not count the point
			break
		}
		dirs, previouslyVisited := pointsVisited[move]
		if previouslyVisited && dirSliceContains(dirs, dir) {
			return true
		}

		_, hitObstacle := obstacles[move]
		if hitObstacle || move == newObstacle {
			dir = nextDirection(dir)
			continue
		}
		guardPosition = move
		pointsVisited[move] = append(pointsVisited[move], dir)
	}

	return false
}

func inBounds(p advent.Point, xOverflow int, yOverflow int) bool {
	return !(p.X < 0 || p.X >= xOverflow || p.Y < 0 || p.Y >= yOverflow)
}

func drawMap(pos advent.Point, pointsVisited map[advent.Point][]advent.Direction, xOverflow int, yOverflow int, obstacles map[advent.Point]struct{}) {
	for y := 0; y < yOverflow; y++ {
		for x := 0; x < xOverflow; x++ {
			p := advent.Point{X: x, Y: y}
			if pos == p {
				fmt.Print("X")
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
	fmt.Printf("pos: %s, points visited: %d", pos, len(pointsVisited))
}

func nextMove(pos advent.Point, dir advent.Direction) advent.Point {
	switch dir {
	case advent.Up:
		return advent.Point{X: pos.X, Y: pos.Y - 1}
	case advent.Right:
		return advent.Point{X: pos.X + 1, Y: pos.Y}
	case advent.Down:
		return advent.Point{X: pos.X, Y: pos.Y + 1}
	case advent.Left:
		return advent.Point{X: pos.X - 1, Y: pos.Y}
	}
	panic("invalid direction")
}

func nextDirection(dir advent.Direction) advent.Direction {
	switch dir {
	case advent.Up:
		return advent.Right
	case advent.Right:
		return advent.Down
	case advent.Down:
		return advent.Left
	case advent.Left:
		return advent.Up
	}
	panic("invalid direction")
}

func dirSliceContains(sl []advent.Direction, val advent.Direction) bool {
	for i := range sl {
		if sl[i] == val {
			return true
		}
	}
	return false
}
