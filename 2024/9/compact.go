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

type fileOrSpace struct {
	id     int
	isFile bool
	length int
}

func main() {
	debug := advent.DebugEnabled()

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("expected one line")
	}
	line := strings.Split(scanner.Text(), "")
	fileOrSpaces := make([]fileOrSpace, len(line))
	for i, chr := range line {
		length, err := strconv.Atoi(chr)
		if err != nil {
			log.Fatalf("not an int %s: %v", chr, err)
		}

		id := 0
		isFile := i%2 == 0
		if isFile {
			id = i / 2
		}

		fileOrSpaces[i] = fileOrSpace{
			id: id,
			// even character positions are files; odd character positions are free space
			isFile: isFile,
			length: length,
		}
	}

	blockLength := 0
	for _, fos := range fileOrSpaces {
		blockLength += fos.length
	}
	blocks := make([]*fileOrSpace, blockLength)

	blockPosition := 0
	for i := range fileOrSpaces {
		if fileOrSpaces[i].isFile {
			for j := 0; j < fileOrSpaces[i].length; j++ {
				blocks[blockPosition+j] = &fileOrSpaces[i]
			}
		}
		blockPosition += fileOrSpaces[i].length
	}

	for {
		if !swap(blocks) {
			break
		}
	}

	if debug {
		printBlocks(blocks)
	}

	fmt.Printf("part 1: %d\n", checksum(blocks))

	// Use the starting blocks to build an array of disk blocks with the total length. Each element is a pointer to a file, or nil
	// iterate over the blocks to populate the initial elements
	// blocks are fully compacted only if a nil element is followed only by other nil elements
	// to compact a block, take the last non-nil element (pointer to a file) and swap it with the first nil element
	// calculate a checksum by iterating over the list
}

func isCompact(blocks []*fileOrSpace) bool {
	foundNil := false
	for i := range blocks {
		if blocks[i] != nil && foundNil {
			return false
		}
		if blocks[i] == nil {
			foundNil = true
		}
	}

	return true
}

func swap(blocks []*fileOrSpace) bool {
	lastFilePos := len(blocks) - 1
	for i := lastFilePos; i >= 0; i-- {
		if blocks[i] != nil {
			lastFilePos = i
			break
		}
	}

	firstNilPos := 0
	for i := firstNilPos; i < len(blocks); i++ {
		if blocks[i] == nil {
			firstNilPos = i
			break
		}
	}

	if firstNilPos < lastFilePos {
		blocks[firstNilPos], blocks[lastFilePos] = blocks[lastFilePos], nil
		log.Printf("swapped positions %d and %d for file ID %d", firstNilPos, lastFilePos, blocks[firstNilPos].id)
		return true
	}
	log.Printf("compacted: firstNilPos %d, lastFilePos: %d", firstNilPos, lastFilePos)
	return false
}

func checksum(blocks []*fileOrSpace) int {
	if !isCompact(blocks) {
		//TODO: log the offending positions
		log.Fatal("not compacted")
	}
	sum := 0
	for i := range blocks {
		if blocks[i] == nil {
			continue
		}
		sum += blocks[i].id * i
	}
	return sum
}

func printBlocks(blocks []*fileOrSpace) {
	// each line represents a block, with "." representing empty space
	for i := range blocks {
		if blocks[i] == nil {
			fmt.Println(".")
		} else {
			fmt.Printf("%d\n", blocks[i].id)
		}
	}
}
