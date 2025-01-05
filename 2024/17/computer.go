package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	c := mustParseInput(os.Stdin)
	log.Printf("%v", c)
}

type opcode int

const (
	opAdv opcode = iota
	opBxl
	opBst
	opJnz
	opBxc
	opOut
	opBdv
	opCdv
)

type computer struct {
	regA, regB, regC int
	program          []int
}

var (
	registerPattern = regexp.MustCompile(`Register (.): (\d+)`)
	programPattern  = regexp.MustCompile(`Program: ([\d,]+)`)
)

func mustParseInput(r io.Reader) computer {
	c := computer{}
	scanner := bufio.NewScanner(r)

	c.regA = mustParseRegisterVal(scanner)
	c.regB = mustParseRegisterVal(scanner)
	c.regC = mustParseRegisterVal(scanner)

	if !scanner.Scan() || len(scanner.Text()) != 0 {
		log.Fatalf("missing empty line between Register values and Program")
	}

	if !scanner.Scan() {
		log.Fatal("missing Program line")
	}
	matches := programPattern.FindStringSubmatch(scanner.Text())
	if len(matches) != 2 {
		log.Fatalf("invalid Program matches %v from %s", matches, scanner.Text())
	}

	programStrs := strings.Split(matches[1], ",")
	program := make([]int, len(programStrs))
	for i := range programStrs {
		program[i] = advent.MustParseInt(programStrs[i])
	}
	c.program = program
	return c
}

func mustParseRegisterVal(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		log.Fatalf("missing Register line")
	}
	matches := registerPattern.FindStringSubmatch(scanner.Text())
	if len(matches) != 3 {
		log.Fatalf("invalid Register matches: %v from %s", matches, scanner.Text())
	}
	return advent.MustParseInt(matches[2])
}
