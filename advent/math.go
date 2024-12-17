package advent

func Abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func GCD(a, b int) int {
	greater, lesser := a, b
	if lesser > greater {
		lesser, greater = greater, lesser
	}

	for {
		q := greater / lesser
		r := greater - q*lesser
		if r == 0 {
			return lesser
		}
		greater = lesser
		lesser = r
	}
}
