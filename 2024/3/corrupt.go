package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/andrei-m/aoc/advent"
)

var pattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	submatches := pattern.FindAllStringSubmatch(string(buf), -1)
	sum := 0
	for _, submatch := range submatches {
		if len(submatch) != 3 {
			log.Fatalf("unexpected match length != 3 for: %v", submatch)
		}
		sum += advent.MustParseInt(submatch[1]) * advent.MustParseInt(submatch[2])
	}
	fmt.Printf("part 1: %d\n", sum)

}
