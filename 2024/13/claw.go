package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	/*
		Each scenario provides coefficients for buttons A and B for each axis, e.g. 50a + 28b = 6314 (example is for one axis)
		Look for (a, b) pairs that satisfy both axes. Iterate between 0 and 100 for each potential value.
		Sort the satisfying pairs by token cost (a==3 tokens; b==1 token), ascending. Select the first pair, which is most efficient. Add its token cost to the sum.
	*/
	scenarios := mustParseScenarios()
	log.Printf("parsed %d scenarios", len(scenarios))

	const part1MaxPresses = 100

	var tokens int
	for _, s := range scenarios {
		sols := getSolutions(s, part1MaxPresses)
		if len(sols) > 0 {
			sort.Slice(sols, func(i, j int) bool {
				return sols[i].score() < sols[j].score()
			})
			log.Printf("scenario: %v; solutions: %v: best score: %d", s, sols, sols[0].score())
			tokens += sols[0].score()
		} else {
			log.Printf("scenario: %v; no solution", s)
		}
	}
	fmt.Printf("part 1: %d\n", tokens)

	tokens = 0
	for _, s := range scenarios {
		s.prizeLoc.X += part2PrizeAddend
		s.prizeLoc.Y += part2PrizeAddend
		sols := getSolutions(s, 0)
		if len(sols) > 0 {
			tokens += sols[0].score()
		}
	}
	fmt.Printf("part 2: %d\n", tokens)
}

type solution struct {
	a int
	b int
}

func (s solution) score() int {
	return s.a*3 + s.b
}

func getSolutions(scen scenario, maxPresses int) []solution {
	/*
		Cramer's rule applied to:
		aButtonPresses * aDeltaX + bButtonPresses * bDeltaX = prizeLoc.X
		aButtonPresses * aDeltaY + bButtonPresses * bDeltaY = prizeLoc.Y

		... to solve for 'aButtonPresses' & 'bButtonPresses'

		matrix for determinant (D):
		aDeltaX		bDeltaX
		aDeltaY		bDeltaY

		aDeltaY		bDeltaY
		aDeltaX		bDeltaX

		matrix D-sub-aButtonPresses (replace prizeLoc constants for aButtonPresses column)
		prizeLocX	bDeltaX
		prizeLocY	bDeltaY

		matrix D-sub-bButtonPresses
		aDeltaX		prizeLocX
		aDeltaY		prizeLocY
	*/
	solutions := []solution{}

	det := scen.aDelta.X*scen.bDelta.Y - scen.bDelta.X*scen.aDelta.Y
	if det == 0 {
		// no solutions or multiple solutions
		//TODO: the problem statement implies multiple solutions are possible, but part 1 input did not reveal any
		return solutions
	}
	detSubA := scen.prizeLoc.X*scen.bDelta.Y - scen.bDelta.X*scen.prizeLoc.Y
	detSubB := scen.aDelta.X*scen.prizeLoc.Y - scen.prizeLoc.X*scen.aDelta.Y

	if detSubA%det != 0 || detSubB%det != 0 {
		// fractional presses are impossible
		return solutions
	}

	a := detSubA / det
	b := detSubB / det
	if a < 0 || b < 0 {
		// negative presses are impossible
		return solutions
	}

	if maxPresses > 0 && (a > maxPresses || b > maxPresses) {
		// discard solutions that exceed the max allowable presses (only set for part 1)
		return solutions
	}

	solutions = append(solutions, solution{a: detSubA / det, b: detSubB / det})
	return solutions
}

type scenario struct {
	aDelta   advent.Point
	bDelta   advent.Point
	prizeLoc advent.Point
}

func mustParseScenarios() []scenario {
	scenarios := []scenario{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		lines, success := getNLines(scanner, 3)
		if !success {
			break
		}
		scen := mustParseScenario(lines)
		//log.Printf("scenario: %v", scen)
		scenarios = append(scenarios, scen)

		// two newlines between scenarios, except for the last
		if !scanner.Scan() {
			break
		}
	}
	return scenarios
}

func getNLines(scanner *bufio.Scanner, n int) ([]string, bool) {
	lines := make([]string, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			if i != 0 {
				// the last scenario of the input does not include a trailing blank line
				log.Fatalf("failed to scan %d lines: EOF on scan #%d", n, i+1)
			}
			return nil, false
		}
		lines[i] = scanner.Text()
	}
	return lines, true
}

var (
	aRegex     = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	bRegex     = regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	prizeRegex = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
)

const part2PrizeAddend = 10000000000000

func mustParseScenario(lines []string) scenario {
	if len(lines) != 3 {
		log.Fatalf("expected 3 lines; got %d", len(lines))
	}
	aSubmatch := aRegex.FindStringSubmatch(lines[0])
	if len(aSubmatch) != 3 {
		log.Fatalf("Button A regexp expected 3 submatches, got %d on line %s", len(aSubmatch), lines[0])
	}
	bSubmatch := bRegex.FindStringSubmatch(lines[1])
	if len(bSubmatch) != 3 {
		log.Fatalf("Button B regexp expected 3 submatches, got %d on line %s", len(bSubmatch), lines[1])
	}
	prizeSubmatch := prizeRegex.FindStringSubmatch(lines[2])
	if len(prizeSubmatch) != 3 {
		log.Fatalf("Price regexp expected 3 submatches, got %d on line %s", len(prizeSubmatch), lines[2])
	}
	return scenario{
		aDelta:   advent.Point{X: advent.MustParseInt(aSubmatch[1]), Y: advent.MustParseInt(aSubmatch[2])},
		bDelta:   advent.Point{X: advent.MustParseInt(bSubmatch[1]), Y: advent.MustParseInt(bSubmatch[2])},
		prizeLoc: advent.Point{X: advent.MustParseInt(prizeSubmatch[1]), Y: advent.MustParseInt(prizeSubmatch[2])},
	}
}
