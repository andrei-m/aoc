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

	for i := 0; i < 25; i++ {
		log.Printf("%v", elements)
		elements = evolve(elements)
		//time.Sleep(5 * time.Second)
	}
	fmt.Printf("part 1: %d\n", len(elements))
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
