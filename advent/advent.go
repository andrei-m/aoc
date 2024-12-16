package advent

import (
	"fmt"
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

func Abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func IntSliceContains(sl []int, val int) bool {
	for i := range sl {
		if sl[i] == val {
			return true
		}
	}
	return false
}

type Point struct {
	X int
	Y int
}

func (p Point) Inbounds(xOverflow, yOverflow int) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < xOverflow && p.Y < yOverflow
}

type Vector struct {
	A Point
	B Point
}

func InvertVector(v Vector) Vector {
	dX := v.A.X - v.B.X
	dY := v.A.Y - v.B.Y
	return Vector{
		A: v.A,
		B: Point{X: v.A.X + dX, Y: v.A.Y + dY},
	}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func DebugEnabled() bool {
	return len(os.Getenv("DEBUG")) > 0
}

type Pair[V any] struct {
	A V
	B V
}
