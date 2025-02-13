package solution

func hammingDistance(x int, y int) int {
	xBin := ""
	yBin := ""

	for x > 0 {
		xBin += string(rune(x%2) + 48)
		x /= 2
	}

	for y > 0 {
		yBin += string(rune(y%2) + 48)
		y /= 2
	}

	difference := 0

	minLength := 0
	maxLength := 0
	maxBin := ""

	if len(xBin) < len(yBin) {
		minLength = len(xBin)
		maxLength = len(yBin)
		maxBin = yBin
	} else {
		minLength = len(yBin)
		maxLength = len(xBin)
		maxBin = xBin
	}

	for i := 0; i < minLength; i++ {
		if xBin[i] != yBin[i] {
			difference++
		}
	}

	for i := minLength; i < maxLength; i++ {
		if maxBin[i] == '1' {
			difference++
		}
	}

	return difference
}
