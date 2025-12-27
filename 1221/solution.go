package solution

func balancedStringSplit(s string) int {
	count := 0
	hasL := 0
	hasR := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			hasL++
		} else {
			hasR++
		}

		if hasR == hasL {
			count++
			hasR = 0
			hasL = 0
		}
	}

	return count
}
