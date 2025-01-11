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
	part1(&c)
	part2(&c)
	/*
		Part 2 thoughts:
		- B & C are derived from A at the start of the program, so there is no need to record their states.
		- The program halts when A is zero, so smaller values of A produce output values that are at the tail of the program's output.
		- Starting with 0, and working up: try a value of A. Also pick a desired output digit, starting with the last digit of the program.
			- Record what digit that value of A outputs
			- If the digit is not the target digit, discard the value of A & try the next one
			- If it is the target digit, continue to the previous target digit
			- When the instruction pointer is 0, check if the value of A was previously observed to output to output a value other than the one desired. If so, discard that value of A
	*/
}

func part1(c *computer) {
	for !c.halt() {
		c.advance()
	}
	fmt.Println("part 1:")
	c.print()
}

func part2(c *computer) {
	regA := 0
	for {
		if tryRegA(c, regA) {
			fmt.Println("part 2:")
			c.print()
			break
		} else {
			if len(c.out) > 8 {
				log.Printf("%d: %v", regA, c.out)
			}
		}
		regA++
	}
}

func tryRegA(c *computer, regA int) bool {
	c.reset(regA)
	outLen := 0

	for !c.halt() {
		c.advance()
		if len(c.out) > outLen {
			// a new value was outputted
			outLen++
			if !c.checkQuinePrefix() {
				// try next regA if output is not a prefix of the program (not a quine)
				return false
			}
			if len(c.out) == len(c.program) {
				// found a prefix that is the full length of the program
				return true
			}
		}
	}

	return false
}

func (c *computer) checkQuinePrefix() bool {
	for i := range c.out {
		if c.out[i] != c.program[i] {
			return false
		}
	}
	return true
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

type state struct {
	regA, regB, regC   int
	instructionPointer int
}

type computer struct {
	state
	program []int
	out     []int
}

func (c computer) String() string {
	return fmt.Sprintf("A: %d B: %d C: %d instruction: %d\nout:%v\n%s",
		c.regA, c.regB, c.regC, c.instructionPointer, c.out, c.nextInstruction())
}

func (c computer) nextInstruction() string {
	sb := strings.Builder{}
	if c.halt() {
		return sb.String()
	}
	sb.WriteString(fmt.Sprintf("next op: %d", c.program[c.instructionPointer]))
	if c.instructionPointer < len(c.program)-1 {
		sb.WriteString(fmt.Sprintf(" combo: %d", c.program[c.instructionPointer+1]))
	}
	return sb.String()
}

func (c computer) print() {
	sb := strings.Builder{}
	for i, val := range c.out {
		if i != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%d", val))
	}
	fmt.Printf("%s\n", sb.String())
}

func (c computer) halt() bool {
	return c.instructionPointer >= len(c.program)
}

func (c *computer) advance() {
	opInt := c.program[c.instructionPointer]
	if opInt < 0 || opInt > 7 {
		log.Fatalf("invalid opcode: %d", opInt)
	}

	previousState := c.state
	//log.Println(c.String())

	op := opcode(opInt)
	switch op {
	case opAdv:
		// A = A / 2^combo
		combo := c.comboVal()
		c.regA = c.regA / (1 << combo)
		c.instructionPointer += 2
	case opBxl:
		// B = B xor literal
		lit := c.program[c.instructionPointer+1]
		c.regB = int(uint(c.regB) ^ uint(lit))
		c.instructionPointer += 2
	case opBst:
		// B = combo % 8
		combo := c.comboVal()
		c.regB = combo % 8
		c.instructionPointer += 2
	case opJnz:
		// jump to literal if A != 0
		if c.regA != 0 {
			c.instructionPointer = c.program[c.instructionPointer+1]
		} else {
			c.instructionPointer += 2
		}
	case opBxc:
		// B = B xor C
		c.regB = int(uint(c.regB) ^ uint(c.regC))
		c.instructionPointer += 2
	case opOut:
		// output mod 8
		val := c.comboVal() % 8
		c.out = append(c.out, val)
		c.instructionPointer += 2
	case opBdv:
		// B = A / 2^combo
		combo := c.comboVal()
		c.regB = c.regA / (1 << combo)
		c.instructionPointer += 2
	case opCdv:
		// C = A / 2^combo
		combo := c.comboVal()
		c.regC = c.regA / (1 << combo)
		c.instructionPointer += 2
	default:
		log.Fatalf("invalid op: %v", op)
	}

	if c.state == previousState {
		log.Fatal("infinite loop")
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

func (c *computer) reset(regA int) {
	c.instructionPointer = 0
	c.regA = regA
	c.regB = 0
	c.regC = 0
	c.out = []int{}
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
