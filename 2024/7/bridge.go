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

type equation struct {
	sum      int
	operands []int
}

func main() {
	equations := []equation{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) != 2 {
			log.Fatalf("invalid line: %v, expected two colon-separated parts", parts)
		}

		sum := advent.MustParseInt(strings.TrimSpace(parts[0]))
		rawOperands := strings.Split(strings.TrimSpace(parts[1]), " ")
		operands := make([]int, len(rawOperands))
		for i := range rawOperands {
			operands[i] = advent.MustParseInt(rawOperands[i])
		}

		equations = append(equations, equation{sum: sum, operands: operands})
	}

	sum := 0
	for _, eq := range equations {
		if satisfiable(eq) {
			sum += eq.sum
		}
	}
	fmt.Printf("part 1: %d\n", sum)
}

type op int

const (
	add op = iota
	mul
	concat
)

func satisfiable(eq equation) bool {
	ops := make([]op, len(eq.operands)-1)

	for i := 0; i < 1<<len(ops); i++ {
		for j := 0; j < len(ops); j++ {
			testBit := 1 << j
			if i&testBit == testBit {
				ops[j] = mul
			} else {
				ops[j] = add
			}
		}

		if eval(eq, ops) == eq.sum {
			printEquation(eq, ops)
			return true
		}
	}

	printUnsatisfiableEquation(eq)
	return false
}

func eval(eq equation, ops []op) int {
	if len(eq.operands) < 2 {
		log.Fatalf("invalid equation: %v", eq)
	}
	if len(ops) != len(eq.operands)-1 {
		log.Fatalf("ops must be 1 fewer than operands")
	}

	result := eq.operands[0]
	for i, op := range ops {
		if op == add {
			result += eq.operands[i+1]
		} else if op == mul {
			result = result * eq.operands[i+1]
		} else if op == concat {
			concatenated := fmt.Sprintf("%d%d", result, eq.operands[i+1])
			concatenatedInt, err := strconv.Atoi(concatenated)
			if err != nil {
				panic(err)
			}
			result = concatenatedInt
		} else {
			panic("invalid op is either 'add' nor 'mul'")
		}
	}
	return result
}

func printUnsatisfiableEquation(eq equation) {
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(eq.sum))
	sb.WriteString(": ")
	for i := range eq.operands {
		sb.WriteString(strconv.Itoa(eq.operands[i]))
		sb.WriteString(" ")
	}
	fmt.Println(sb.String())
}

func printEquation(eq equation, ops []op) {
	if len(eq.operands) < 2 {
		log.Fatalf("invalid equation: %v", eq)
	}
	if len(ops) != len(eq.operands)-1 {
		log.Fatalf("ops must be 1 fewer than operands")
	}

	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(eq.sum))
	sb.WriteString(": ")
	sb.WriteString(strconv.Itoa(eq.operands[0]))

	for i, op := range ops {
		if op == add {
			sb.WriteString("+")
		} else if op == mul {
			sb.WriteString("*")
		} else if op == concat {
			sb.WriteString("||")
		} else {
			panic("invalid operator")
		}
		sb.WriteString(strconv.Itoa(eq.operands[i+1]))
	}

	fmt.Println(sb.String())
}
