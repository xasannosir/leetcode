package solution

func isAnagram(s string, t string) bool {
	if len(s) > len(t) || len(t) > len(s) {
		return false
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				t = t[:j] + t[j+1:]
				break
			}
		}
	}

	return t == ""
}
