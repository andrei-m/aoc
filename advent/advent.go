package advent

import (
	"log"
	"strconv"
)

func MustParseInt(raw string) int {
	leftInt, err := strconv.Atoi(raw)
	if err != nil {
		log.Fatalf("not an int: %s", raw)
	}
	return leftInt
}

func Abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
