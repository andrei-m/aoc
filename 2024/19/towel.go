package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	/*
		part 1:
		- Sequences composed of single stripe colors that are available are easier to satisfy, so prioritize longer patterns
		- For a given sequence, identify the longest towels that can fit.
		- For each fitting towel, place it to create a matched subsequence & sub-sequences that require matches.
		- For each subsequence, recursively attempt to identify the longest towel that can fit.
		- A sequence is matched if a fitting towel is found & there are no sub-sequences, or every sub-sequence is recursively matched
	*/
	towels, patterns := mustParseInput(os.Stdin)
	fmt.Printf("part 1: %d\n", part1(towels, patterns))
	fmt.Printf("part 2: %d\n", part2(towels, patterns))
}

func part1(towels, patterns []string) int {
	sort.Slice(towels, func(i, j int) bool {
		return len(towels[i]) > len(towels[j])
	})

	memo := map[string]int{}
	possiblePatternCount := 0
	for _, p := range patterns {
		if countPossibilities(memo, towels, p) > 0 {
			log.Printf("possible: %s", p)
			possiblePatternCount++
		} else {
			log.Printf("not possible: %s", p)
		}
	}
	return possiblePatternCount
}

func part2(towels, patterns []string) int {
	sort.Slice(towels, func(i, j int) bool {
		return len(towels[i]) > len(towels[j])
	})

	memo := map[string]int{}
	totalPossibilities := 0
	for _, p := range patterns {
		c := countPossibilities(memo, towels, p)
		log.Printf("%s: %d", p, c)
		totalPossibilities += c
	}
	return totalPossibilities
}

func countPossibilities(memo map[string]int, towels []string, pattern string) int {
	if len(pattern) == 0 {
		// assume an empty pattern means exactly 1 way to satisfy a previous non-empty pattern
		return 1
	}

	c, found := memo[pattern]
	if found {
		return c
	}

	count := 0
	for _, t := range towels {
		if len(t) <= len(pattern) && t == pattern[:len(t)] {
			count += countPossibilities(memo, towels, pattern[len(t):])
		}
	}

	memo[pattern] = count
	return count
}

func mustParseInput(r io.Reader) ([]string, []string) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		log.Fatal("expected a 'towels' line")
	}

	rawTowels := scanner.Text()
	towels := strings.Split(rawTowels, ",")
	for i := range towels {
		towels[i] = strings.TrimSpace(towels[i])
	}

	if !scanner.Scan() || len(scanner.Text()) != 0 {
		log.Fatal("expected an empty separator line between towels and patterns")
	}

	patterns := []string{}
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}

	return towels, patterns
}
