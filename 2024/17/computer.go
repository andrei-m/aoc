package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/andrei-m/aoc/advent"
)

func main() {
	c := mustParseInput(os.Stdin)
	part1(c)
}

func part1(c computer) {
	for !c.halt() {
		c.advance()
	}
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
	regA, regB, regC   int
	program            []int
	instructionPointer int
}

func (c computer) halt() bool {
	return c.instructionPointer >= len(c.program)
}

func (c *computer) advance() {
	opInt := c.program[c.instructionPointer]
	if opInt < 0 || opInt > 7 {
		log.Fatalf("invalid opcode: %d", opInt)
	}
	op := opcode(opInt)
	switch op {
	case opAdv:
		// division
		combo := c.comboVal()
		c.regA = c.regA / (2 << combo)
		c.instructionPointer += 2
	case opBxl:
		// bitwise xor
		combo := c.comboVal()
		c.regB = c.regB ^ combo
		c.instructionPointer += 2
	case opBst:
		// combo mod 8
		combo := c.comboVal()
		c.regB = combo % 8
		c.instructionPointer += 2
	case opJnz:
		// jump
		if c.regA != 0 {
			c.instructionPointer = c.comboVal()
		}
	case opBxc:
		// bitwise xor with registers B+C
		c.regB = c.regB ^ c.regC
		c.instructionPointer += 2
	case opOut:
		// output mod 8
		val := c.comboVal() % 8
		fmt.Printf("%d", val)
		c.instructionPointer += 2
	case opBdv:
		// division stored to register B
		combo := c.comboVal()
		c.regB = c.regA / (2 << combo)
		c.instructionPointer += 2
	case opCdv:
		// division stored to register C
		combo := c.comboVal()
		c.regC = c.regA / (2 << combo)
		c.instructionPointer += 2
	default:
		log.Fatalf("invalid op: %v", op)
	}
}

func (c computer) comboVal() int {
	combo := c.program[c.instructionPointer+1]
	switch combo {
	case 0, 1, 2, 3:
		return combo
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	default:
		log.Fatalf("invalid combo value: %d", combo)
	}
	panic("unreachable")
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
