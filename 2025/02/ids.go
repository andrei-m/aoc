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

	p2Sum := 0
	for i := range ranges {
		invalidIDs := getInvalidIDsP2(ranges[i])
		for j := range invalidIDs {
			p2Sum += invalidIDs[j]
		}
	}
	fmt.Printf("part 2: %d\n", p2Sum)
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

func getInvalidIDsP2(idr idRange) []int {
	invalidIDs := []int{}
	for i := idr.A; i <= idr.B; i++ {
		if !isValid(i) {
			invalidIDs = append(invalidIDs, i)
		}
	}
	return invalidIDs
}

func isValid(val int) bool {
	// iterate through each sub-sequence from index 0 through len(i)/2 (longer sequences have insufficient space to repeat)
	// Check if the next len(seq) characters repeat. If so, continue checking subsequent len(seq) characters.
	// If at least one repetition is found & end of the string is reached, the id is invalid.
	// If the string has insufficient characters or a non-repetition is encountered, proceed to the next sequence.
	valStr := strconv.Itoa(val)
	for i := 0; i < len(valStr)/2; i++ {
		seq := valStr[0 : i+1]
		remainder := valStr[len(seq):]
		if len(remainder) == 0 {
			return true
		}
		if remainderRepeatsSequence(seq, remainder) {
			return false
		}
	}
	return true
}

func remainderRepeatsSequence(seq, remainder string) bool {
	if len(remainder) == 0 {
		return false
	}
	if len(remainder) < len(seq) {
		return false
	}
	if remainder[0:len(seq)] == seq {
		newRemainder := remainder[len(seq):]
		if len(newRemainder) == 0 {
			return true
		}
		return remainderRepeatsSequence(seq, newRemainder)
	}
	return false
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
