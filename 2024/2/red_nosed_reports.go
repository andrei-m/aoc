package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
			safeDampenedCount++
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
	fmt.Printf("part 2: %d\n", safeDampenedCount)
}

func isSafe(levels []string) (bool, error) {
	if len(levels) == 0 {
		return true, nil
	}
	previous := mustGetInt(levels[0])
	previousIncreasing := false

	for i := 1; i < len(levels); i++ {
		current := mustGetInt(levels[i])

		diff := previous - current
		absDiff := abs(diff)
		if absDiff < 1 || absDiff > 3 {
			//log.Printf("absDiff index %d: %d", i, absDiff)
			return false, nil
		}
		currentIncreasing := diff < 0

		if i > 1 && (currentIncreasing != previousIncreasing) {
			//log.Printf("curentIncreasing: %t, previousIncreasting: %t", currentIncreasing, previousIncreasing)
			return false, nil
		}

		previous = mustGetInt(levels[i])
		previousIncreasing = currentIncreasing
	}

	return true, nil
}

func mustGetInt(raw string) int {
	val, err := strconv.Atoi(raw)
	if err != nil {
		log.Fatalf("not an int %s: %v", raw, err)
	}
	return val
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
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
