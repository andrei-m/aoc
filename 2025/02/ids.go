package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	ranges := mustParseInput(os.Stdin)
	log.Printf("parsed %d ranges", len(ranges))

	p1Sum := 0
	for i := range ranges {
		invalidIDs := getInvalidIDs(ranges[i])
		for j := range invalidIDs {
			p1Sum += invalidIDs[j]
		}
	}
	fmt.Printf("part 1: %d\n", p1Sum)
}

func getInvalidIDs(idr idRange) []int {
	invalidIDs := []int{}
	for i := idr.A; i <= idr.B; i++ {
		iStr := strconv.Itoa(i)
		if len(iStr)%2 != 0 {
			// odd lengths are all valid
			continue
		}
		mid := len(iStr) / 2
		leftHalf := iStr[0:mid]
		rightHalf := iStr[mid:]
		if leftHalf == rightHalf {
			invalidIDs = append(invalidIDs, i)
		}
	}
	return invalidIDs
}

type idRange advent.Pair[int]

func mustParseInput(r io.Reader) []idRange {
	scanner := bufio.NewScanner(r)
	ranges := []idRange{}
	if !scanner.Scan() {
		log.Fatal("expected one line")
	}
	rangeStrings := strings.Split(scanner.Text(), ",")
	for i := range rangeStrings {
		parts := strings.Split(rangeStrings[i], "-")
		if len(parts) != 2 {
			log.Fatalf("corrupt range at index %d: %q", i, rangeStrings[i])
		}
		low, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("corrupt low part of range %q: %v", parts[0], err)
		}
		high, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("corrupt high part of range %q: %v", parts[1], err)
		}
		ranges = append(ranges, idRange{A: low, B: high})
	}
	return ranges
}
