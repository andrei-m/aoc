package advent

import "fmt"

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

// TODO: test me
func AddVector(v1, v2 Vector) Vector {
	dX := v2.B.X - v2.A.X
	dY := v2.B.Y - v2.A.Y
	return Vector{
		A: v1.A,
		B: Point{X: v1.B.X + dX, Y: v1.B.Y + dY},
	}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func AdjacentDirFn(xOverflow, yOverflow int) func(Point, Direction) *Point {
	return func(loc Point, dir Direction) *Point {
		switch dir {
		case Left:
			if loc.X <= 0 {
				return nil
			}
			return &Point{X: loc.X - 1, Y: loc.Y}
		case Right:
			if loc.X >= xOverflow {
				return nil
			}
			return &Point{X: loc.X + 1, Y: loc.Y}
		case Up:
			if loc.Y <= 0 {
				return nil
			}
			return &Point{X: loc.X, Y: loc.Y - 1}
		case Down:
			if loc.Y >= yOverflow {
				return nil
			}
			return &Point{X: loc.X, Y: loc.Y + 1}
		default:
			panic("invalid direction")
		}
	}
}

func AdjacentsFn(xOverflow, yOverflow int) func(Point) []Point {
	return func(loc Point) []Point {
		adj := []Point{}
		if loc.X > 0 {
			// left
			adj = append(adj, Point{X: loc.X - 1, Y: loc.Y})
		}
		if loc.X < xOverflow-1 {
			// right
			adj = append(adj, Point{X: loc.X + 1, Y: loc.Y})
		}
		if loc.Y > 0 {
			// up
			adj = append(adj, Point{X: loc.X, Y: loc.Y - 1})
		}
		if loc.Y < yOverflow-1 {
			// down
			adj = append(adj, Point{X: loc.X, Y: loc.Y + 1})
		}
		return adj
	}
}
