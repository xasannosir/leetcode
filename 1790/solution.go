package solution

func areAlmostEqual(s1 string, s2 string) bool {
	i1 := -1
	i2 := -1

	for i := range s1 {
		if s1[i] == s2[i] {
			continue
		}

		if i1 == -1 {
			i1 = i
		} else if i2 == -1 {
			i2 = i
		} else {
			return false
		}
	}

	if i1 == -1 {
		return true
	}

	if i2 == -1 {
		return false
	}

	return s1[i1] == s2[i2] && s1[i2] == s2[i1]
}
