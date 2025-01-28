package solution

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sPattern, tPattern := map[byte]int{}, map[byte]int{}

	for index := range s {
		if sPattern[s[index]] != tPattern[t[index]] {
			return false
		} else {
			sPattern[s[index]] = index + 1
			tPattern[t[index]] = index + 1
		}
	}

	return true
}
