package advent

import (
	"log"
	"os"
	"strconv"
)

func MustParseInt(raw string) int {
	leftInt, err := strconv.Atoi(raw)
	if err != nil {
		log.Fatalf("not an int: %s", raw)
	}
	return leftInt
}

func IntSliceContains(sl []int, val int) bool {
	for i := range sl {
		if sl[i] == val {
			return true
		}
	}
	return false
}

func DebugEnabled() bool {
	return len(os.Getenv("DEBUG")) > 0
}

type Pair[V any] struct {
	A V
	B V
}
