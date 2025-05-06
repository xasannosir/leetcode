package solution

func rotateString(s string, goal string) bool {
	for i := 0; i < len(s); i++ {
		if s[i:]+s[:i] == goal {
			return true
		}
	}

	return false
}
