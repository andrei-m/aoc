package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("expected one line")
	}

	elements := strings.Split(scanner.Text(), " ")
	orig := make([]string, len(elements))
	copy(orig, elements)

	for i := 0; i < 25; i++ {
		//log.Printf("%v", elements)
		elements = evolve(elements)
	}
	fmt.Printf("part 1: %d\n", len(elements))

	elementsMap := make(map[int]int, len(orig))
	for _, elem := range orig {
		elementsMap[advent.MustParseInt(elem)] = 1
	}
	for i := 0; i < 75; i++ {
		log.Printf("iteration %d: len(elements): %d", i+1, score(elementsMap))
		elementsMap = evolvePart2(elementsMap)
	}

	fmt.Printf("part 2: %d\n", score(elementsMap))
}

func score(elements map[int]int) int {
	sum := 0
	for _, v := range elements {
		sum += v
	}
	return sum
}

func evolve(elements []string) []string {
	next := []string{}
	for i := range elements {
		if elements[i] == "0" {
			next = append(next, "1")
		} else if len(elements[i])%2 == 0 {
			next = append(next, split(elements[i])...)
		} else {
			next = append(next, mult(elements[i]))
		}
	}
	return next
}

func evolvePart2(elements map[int]int) map[int]int {
	next := map[int]int{}
	for k := range elements {
		if k == 0 {
			next[1] += elements[k]
		} else if evenDigits(k) {
			split := splitInt(k)
			next[split[0]] += elements[k]
			next[split[1]] += elements[k]
		} else {
			next[k*2024] += elements[k]
		}
	}
	return next
}

func split(elem string) []string {
	first := elem[:len(elem)/2]
	last := elem[len(elem)/2:]
	firstInt := advent.MustParseInt(first)
	lastInt := advent.MustParseInt(last)
	return []string{strconv.Itoa(firstInt), strconv.Itoa(lastInt)}
}

func mult(elem string) string {
	return strconv.Itoa(advent.MustParseInt(elem) * 2024)
}

var evenDigitsMemo = map[int]bool{}

func evenDigits(elem int) bool {
	even, ok := evenDigitsMemo[elem]
	if ok {
		return even
	}
	str := strconv.Itoa(elem)
	even = len(str)%2 == 0
	evenDigitsMemo[elem] = even
	return even
}

var splitIntMemo = map[int][]int{}

func splitInt(elem int) []int {
	split, ok := splitIntMemo[elem]
	if ok {
		return split
	}
	str := strconv.Itoa(elem)
	first := str[:len(str)/2]
	last := str[len(str)/2:]
	firstInt := advent.MustParseInt(first)
	lastInt := advent.MustParseInt(last)
	splitIntMemo[elem] = []int{firstInt, lastInt}
	return splitIntMemo[elem]
}
