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

	// Part 1:
	// Use the starting blocks to build an array of disk blocks with the total length. Each element is a pointer to a file, or nil
	// iterate over the blocks to populate the initial elements
	// blocks are fully compacted only if a nil element is followed only by other nil elements
	// to compact a block, take the last non-nil element (pointer to a file) and swap it with the first nil element
	// calculate a checksum by iterating over the list
	blocks := createBlocks(fileOrSpaces)
	for {
		if !swap(blocks) {
			break
		}
	}
	/*
		if debug {
			printBlocks(blocks)
		}
	*/
	fmt.Printf("part 1: %d\n", checksum(blocks))

	blocks = createBlocks(fileOrSpaces)
	swapAttempted := map[int]struct{}{}
	for {
		if !swapPart2(blocks, swapAttempted) {
			break
		}
	}

	if debug {
		printBlocks(blocks)
	}
	fmt.Printf("part 2: %d\n", checksum(blocks))
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

func swapPart2(blocks []*fileOrSpace, moveAttempted map[int]struct{}) bool {
	// Find start of last file that has not already been moved; if no such file, return false
	lastFileStartPos := len(blocks) - 1
	var lastFile *fileOrSpace
	for i := lastFileStartPos; i >= 0; i-- {
		if blocks[i] != nil {
			_, attempted := moveAttempted[blocks[i].id]
			if attempted {
				continue
			}

			lastFileStartPos = i - (blocks[i].length - 1)
			lastFile = blocks[i]
			break
		}
	}

	// already attempted to move all files
	if lastFile == nil {
		return false
	}
	moveAttempted[lastFile.id] = struct{}{}

	// find first nil position of a nil sequence of at least the length of that file; if no such sequence, return false
	freeBlockStartPos := len(blocks)
	for i := range blocks {
		if blocks[i] != nil {
			continue
		}
		nextBlocksOK := true
		for j := 1; j < lastFile.length; j++ {
			if (i+j) >= len(blocks) || blocks[i+j] != nil {
				nextBlocksOK = false
				break
			}
		}

		if nextBlocksOK {
			freeBlockStartPos = i
			break
		}
	}

	if freeBlockStartPos == len(blocks) {
		// no free space, but try other files
		return true
	}
	if freeBlockStartPos > lastFileStartPos {
		// free space comes after the file; don't swap, but try other files
		return true
	}

	// swap n positions for file length 'n'
	for i := 0; i < lastFile.length; i++ {
		blocks[freeBlockStartPos+i], blocks[lastFileStartPos+i] = blocks[lastFileStartPos+i], nil
	}
	log.Printf("swapped %d nil blocks starting at %d with file ID %d starting at %d", lastFile.length, freeBlockStartPos, lastFile.id, lastFileStartPos)
	return true
}

func isCompactPart2(blocks []*fileOrSpace) bool {
	for i := range blocks {
		if blocks[i] == nil {
			continue
		}

		consecutiveNils := 0
		for j := 0; j < i; j++ {
			if blocks[j] != nil {
				consecutiveNils = 0
				// reset nil segment
				continue
			}
			consecutiveNils++
			if consecutiveNils >= blocks[i].length {
				// found earlier segment of nils
				log.Printf("found %d nils starting at position %d for file ID %d", consecutiveNils, j, blocks[i].id)
				return false
			}
		}
	}
	return true
}

func checksum(blocks []*fileOrSpace) int {
	/*
		        //TODO: This doesn't account for the rule that requires a file to move at most once because it fails if there is any continguous free space to the left of the file (even if that space was freed after the file moved)
				if !isCompactPart2(blocks) {
					log.Fatal("not compact")
				}
	*/
	sum := 0
	for i := range blocks {
		if blocks[i] == nil {
			continue
		}
		sum += blocks[i].id * i
	}
	return sum
}

func createBlocks(fileOrSpaces []fileOrSpace) []*fileOrSpace {
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
	return blocks
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
