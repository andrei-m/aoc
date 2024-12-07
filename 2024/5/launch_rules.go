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
	rules := []rulePair{}
	sequences := [][]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		lineSplit := strings.Split(line, "|")
		if len(lineSplit) == 2 {
			first := advent.MustParseInt(lineSplit[0])
			last := advent.MustParseInt(lineSplit[1])
			rules = append(rules, rulePair{first, last})
			continue
		}

		lineSplit = strings.Split(line, ",")
		if len(lineSplit) > 1 {
			seq := make([]int, len(lineSplit))
			for i := range lineSplit {
				seq[i] = advent.MustParseInt(lineSplit[i])
			}
			sequences = append(sequences, seq)
		}
	}

	log.Printf("rule count: %d, sequence count: %d", len(rules), len(sequences))

	ruleIndex := map[int][]int{}
	for _, r := range rules {
		ruleIndex[r.last] = append(ruleIndex[r.last], r.first)
	}

	sum := 0
	for _, seq := range sequences {
		if isOrdered(seq, ruleIndex) {
			sum += seq[len(seq)/2]
		}
	}
	fmt.Printf("sum: %d", sum)
}

// isOrdered returns true if each pair of (element in seq, each proceeding element in seq) does not violate the rules represented in ruleIndex
// ruleIndex is keyed by the "last" element, e.g. the rule 46 -> [91, 22] means 22 and 91 must both preceed 46
func isOrdered(seq []int, ruleIndex map[int][]int) bool {
	for i := range seq {
		for j := 0; j < i; j++ {
			requiredPredecessors := ruleIndex[seq[j]]
			if advent.IntSliceContains(requiredPredecessors, seq[i]) {
				log.Printf("FALSE: %v; %d|%d; rule %d -> %v", seq, seq[i], seq[j], seq[j], ruleIndex[seq[j]])
				return false
			}
		}
	}

	log.Printf("TRUE: %v", seq)
	return true
}

type rulePair struct {
	first int
	last  int
}
