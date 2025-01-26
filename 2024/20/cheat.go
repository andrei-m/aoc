package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	maze := mustParseInput(os.Stdin)
	fmt.Printf("%v", maze)
}

type object int

const (
	path object = iota
	wall
	start
	end
)

func mustParseInput(r io.Reader) [][]object {
	scanner := bufio.NewScanner(r)
	rows := [][]object{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		row := make([]object, len(line))
		for i := range line {
			switch line[i] {
			case ".":
				row[i] = path
			case "#":
				row[i] = wall
			case "S":
				row[i] = start
			case "E":
				row[i] = end
			default:
				log.Fatalf("invalid char: %s", line[i])
			}
		}
		rows = append(rows, row)
	}

	return rows
}
