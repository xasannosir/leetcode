package solution

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for ; i < len(s); i++ {
		isHave := false

		for ; j < len(t); j++ {
			if s[i] == t[j] {
				j++
				isHave = true
				break
			}
		}

		if !isHave {
			return false
		}
	}

	return true
}
