package helpers

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Max4(a, b, c, d int) int {
	return Max(Max(a, b), Max(c, d))
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Min4(a, b, c, d int) int {
	return Min(Min(a, b), Min(c, d))
}
