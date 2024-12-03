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
	safeCount := 0
	safeDampenedCount := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Fields(line)

		safe, err := isSafe(levels)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if safe {
			safeCount++
		}
		log.Printf("%t: %v", safe, levels)

		if !safe {
			for i := range levels {
				iRemoved := removeCopy(levels, i)
				safe, err := isSafe(iRemoved)
				if err != nil {
					log.Fatalf("%v", err)
				}
				log.Printf("%v dampened %t: %v", levels, safe, iRemoved)
				if safe {
					safeDampenedCount++
					break
				}
			}
		}

	}

	fmt.Printf("part 1: %d\n", safeCount)
	fmt.Printf("part 2: %d\n", safeCount+safeDampenedCount)
}

func isSafe(levels []string) (bool, error) {
	if len(levels) == 0 {
		return true, nil
	}
	previous := advent.MustParseInt(levels[0])
	previousIncreasing := false

	for i := 1; i < len(levels); i++ {
		current := advent.MustParseInt(levels[i])

		diff := previous - current
		absDiff := advent.Abs(diff)
		if absDiff < 1 || absDiff > 3 {
			//log.Printf("absDiff index %d: %d", i, absDiff)
			return false, nil
		}
		currentIncreasing := diff < 0

		if i > 1 && (currentIncreasing != previousIncreasing) {
			//log.Printf("curentIncreasing: %t, previousIncreasting: %t", currentIncreasing, previousIncreasing)
			return false, nil
		}

		previous = current
		previousIncreasing = currentIncreasing
	}

	return true, nil
}

func removeCopy(slice []string, i int) []string {
	result := make([]string, 0, len(slice)-1)
	for idx := range slice {
		if idx == i {
			continue
		}
		result = append(result, slice[idx])
	}
	return result
}
