package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
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

		left = append(left, mustParseInt(elements[0]))
		right = append(right, mustParseInt(elements[1]))
	}
	if scanner.Err() != nil {
		log.Fatalf("failed to read input: %v", scanner.Err())
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0
	for i := range left {
		distance += abs(left[i] - right[i])
	}
	fmt.Printf("%d\n", distance)
}

func mustParseInt(raw string) int {
	leftInt, err := strconv.Atoi(raw)
	if err != nil {
		log.Fatalf("not an int: %s", raw)
	}
	return leftInt
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
