package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	left, right := []int{}, []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Fields(line)
		if len(elements) != 2 {
			log.Fatalf("malformed line: %s", line)
		}

		left = append(left, advent.MustParseInt(elements[0]))
		right = append(right, advent.MustParseInt(elements[1]))
	}
	if scanner.Err() != nil {
		log.Fatalf("failed to read input: %v", scanner.Err())
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0
	for i := range left {
		distance += advent.Abs(left[i] - right[i])
	}
	fmt.Printf("part 1; difference: %d\n", distance)

	similarity := 0
	for _, val := range left {
		foundIdx := find(val, right)
		if foundIdx == -1 {
			continue
		}
		count := countOccurrences(foundIdx, right)
		similarity += val * count
	}
	fmt.Printf("part 2; similarity: %d\n", similarity)
}

// find returns the index of the value in sortedList or '-1' if it doesn't exist
func find(val int, sortedList []int) int {
	lower := 0
	upper := len(sortedList) - 1
	pivot := 0

	for {
		newPivot := lower + (upper-lower)/2
		if newPivot == pivot {
			return -1
		}
		pivot = newPivot

		if sortedList[pivot] == val {
			return pivot
		} else if sortedList[pivot] > val {
			// search first half
			upper = pivot
		} else {
			lower = pivot
		}
	}
}

func countOccurrences(idx int, sortedList []int) int {
	count := 1

	// Count earlier same-values
	for i := idx - 1; i > 0; i-- {
		if sortedList[i] != sortedList[idx] {
			break
		}
		count++
	}

	// Count later same-values
	for i := idx + 1; i < len(sortedList); i++ {
		if sortedList[i] != sortedList[idx] {
			break
		}
		count++
	}

	return count
}
