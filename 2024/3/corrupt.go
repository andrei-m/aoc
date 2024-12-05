package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/andrei-m/aoc/advent"
)

var pattern = regexp.MustCompile(`(.*?)?mul\((\d{1,3}),(\d{1,3})\)`)

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	submatches := pattern.FindAllStringSubmatch(string(buf), -1)

	sum := 0
	enabledSum := 0
	enabled := true

	for _, submatch := range submatches {
		if len(submatch) != 4 {
			log.Fatalf("unexpected match length != 4 for: %v: %d", submatch, len(submatch))
		}
		newEnabled := toggleEnabled(submatch[1])
		if newEnabled != nil {
			enabled = *newEnabled
		}
		addition := advent.MustParseInt(submatch[2]) * advent.MustParseInt(submatch[3])
		sum += addition
		if enabled {
			enabledSum += addition
		}
		log.Printf("%v enabled: %t, sum: %d, enabledSum: %d", submatch, enabled, sum, enabledSum)
	}
	fmt.Printf("part 1: %d\n", sum)
	fmt.Printf("part 2: %d\n", enabledSum)
}

var (
	do   = regexp.MustCompile(`do\(\)`)
	dont = regexp.MustCompile(`don't\(\)`)
)

func toggleEnabled(prefix string) *bool {
	var result bool
	doMatches := do.FindAllIndex([]byte(prefix), -1)
	dontMatches := dont.FindAllIndex([]byte(prefix), -1)

	log.Printf("prefix: %s, doMatches: %v, dontMatches: %v", prefix, doMatches, dontMatches)

	if len(doMatches) > 0 && len(dontMatches) > 0 {
		lastDo := doMatches[len(doMatches)-1]
		lastDont := dontMatches[len(dontMatches)-1]
		result = lastDo[0] > lastDont[0]
		return &result
	} else if len(doMatches) > 0 {
		result = true
		return &result
	} else if len(dontMatches) > 0 {
		result = false
		return &result
	}

	return nil
}
