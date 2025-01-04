package advent

type Direction int

func (d Direction) String() string {
	switch d {
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	}
	return "unknown"
}

const (
	Up Direction = iota
	Right
	Down
	Left
)

var Dirs = []Direction{Up, Right, Down, Left}

func RotateCounterClockwise(dir Direction) Direction {
	if dir == Up {
		return Left
	}
	return Direction(int(dir-1) % len(Dirs))
}

func RotateClockwise(dir Direction) Direction {
	return Direction(int(dir+1) % len(Dirs))
}
