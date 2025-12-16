package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	banks := mustParseInput(os.Stdin)
	fmt.Printf("part 1: %d\n", part1(banks))
}

func part1(banks [][]int) int {
	sum := 0
	for i := range banks {
		sum += maxJoltage(banks[i])
	}
	return sum
}

func maxJoltage(bank []int) int {
	highestFirst := 0
	highestSecond := 1

	for i := 1; i < len(bank); i++ {
		// The last digit of a bank is not eligible to be the highest because every two digit value is higher joltage
		if bank[i] > bank[highestFirst] && i < len(bank)-1 {
			highestFirst = i
			highestSecond = i + 1
		} else if bank[i] > bank[highestSecond] {
			highestSecond = i
		}
	}

	joltageStr := fmt.Sprintf("%d%d", bank[highestFirst], bank[highestSecond])
	joltage, err := strconv.Atoi(joltageStr)
	if err != nil {
		panic(fmt.Sprintf("%q is not an int; this should not be possible", joltageStr))
	}
	return joltage
}

func mustParseInput(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	banks := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		bank := make([]int, len(line))
		for i := range len(line) {
			joltage, err := strconv.Atoi(string(line[i]))
			if err != nil {
				log.Fatalf("invalid joltage %q: %v", line[i], err)
			}
			bank[i] = joltage
		}
		banks = append(banks, bank)
	}
	return banks
}
