package solution

func lengthOfLastWord(s string) int {
	var count int

	for {
		if s[len(s)-1] == ' ' {
			s = s[:len(s)-1]
		} else {
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else {
			break
		}
	}

	return count
}
