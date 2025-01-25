package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

var debug = false

func main() {
	debug = advent.DebugEnabled()
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
}

func part1(towels, patterns []string) int {
	sort.Slice(towels, func(i, j int) bool {
		return len(towels[i]) > len(towels[j])
	})

	memo := map[string]bool{}
	possiblePatternCount := 0
	for _, p := range patterns {
		if isPossible(memo, towels, p) {
			log.Printf("%s is possible", p)
			possiblePatternCount++
		}
	}
	return possiblePatternCount
}

func isPossible(memo map[string]bool, towels []string, pattern string) bool {
	possible, found := memo[pattern]
	if found {
		return possible
	}
	if debug {
		log.Println(pattern)
	}

	for _, t := range towels {
		for i := 0; i+len(t) <= len(pattern); i++ {
			if t == pattern[i:i+len(t)] {
				if debug {
					log.Printf("%s matched at index %d", t, i)
				}

				if len(t) == len(pattern) {
					// no remaining patterns to match on either end of the towel
					memo[pattern] = true
					return true
				} else if i > 0 && i+len(t) < len(pattern) {
					// remaining patterns exist on both ends
					if isPossible(memo, towels, pattern[0:i]) && isPossible(memo, towels, pattern[i+len(t):]) {
						memo[pattern] = true
						return true
					}
				} else if i > 0 {
					// remaining pattern exists the left only
					if isPossible(memo, towels, pattern[0:i]) {
						memo[pattern] = true
						return true
					}
				} else {
					// remaining pattern exists on the right only
					if isPossible(memo, towels, pattern[i+len(t):]) {
						memo[pattern] = true
						return true
					}
				}
			}
		}
	}

	if debug {
		log.Printf("%s is not possible", pattern)
	}
	memo[pattern] = false
	return false
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
