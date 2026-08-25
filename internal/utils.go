package internal

func intCloseToZero(a, tolerance int) bool {
	return intsCloseEnough(a, 0, tolerance)
}

func intsCloseEnough(a, b, tolerance int) bool {
	return intAbs(a, b) <= tolerance
}

func intAbs(a, b int) int {
	diff := b - a
	if diff < 0 {
		return -diff
	}

	return diff
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
